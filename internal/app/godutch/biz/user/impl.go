package user

import (
	"time"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/base/encrypt"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/blackhorseya/godutch/internal/pkg/infra/token"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
	token  *token.Factory
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node, token *token.Factory) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "UserBiz")),
		repo:   repo,
		node:   node,
		token:  token,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *user.Profile, err error) {
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetUserByID.Error(), zap.Int64("id", id))
		return nil, er.ErrGetUserByID
	}
	if ret == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrUserNotExists
	}

	return ret, nil
}

func (i *impl) GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	if len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		return nil, er.ErrMissingToken
	}

	claims, err := i.token.ValidateToken(token)
	if err != nil {
		i.logger.Error(er.ErrValidateToken.Error(), zap.Error(err))
		return nil, er.ErrValidateToken
	}

	exists, err := i.repo.GetByID(ctx, claims.ID)
	if err != nil {
		i.logger.Error(er.ErrGetUserByID.Error(), zap.Error(err))
		return nil, er.ErrGetUserByID
	}
	if exists == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.Int64("id", claims.ID))
		return nil, er.ErrUserNotExists
	}

	if exists.Token != token {
		i.logger.Error(er.ErrInvalidToken.Error())
		return nil, er.ErrInvalidToken
	}

	return exists, nil
}

func (i *impl) Signup(ctx contextx.Contextx, email, password, name string) (info *user.Profile, err error) {
	if len(email) == 0 {
		i.logger.Error(er.ErrEmptyEmail.Error())
		return nil, er.ErrEmptyEmail
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrEmptyPassword.Error())
		return nil, er.ErrEmptyPassword
	}

	exists, err := i.repo.GetByEmail(ctx, email)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.String("email", email))
		return nil, er.ErrGetUserByEmail
	}
	if exists != nil {
		i.logger.Error(er.ErrUserEmailExists.Error(), zap.String("email", email))
		return nil, er.ErrUserEmailExists
	}

	salt, err := encrypt.HashAndSalt(password)
	if err != nil {
		return nil, er.ErrEncryptPassword
	}

	ret, err := i.repo.Register(ctx, &user.Profile{
		ID:        i.node.Generate().Int64() / 1000 * 1000,
		Email:     email,
		Password:  salt,
		Name:      name,
		CreatedAt: time.Now().UnixNano(),
	})
	if err != nil {
		i.logger.Error(er.ErrSignup.Error(), zap.String("email", email))
		return nil, er.ErrSignup
	}

	return ret, nil
}

func (i *impl) Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	if len(email) == 0 {
		i.logger.Error(er.ErrEmptyEmail.Error())
		return nil, er.ErrEmptyEmail
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrEmptyPassword.Error())
		return nil, er.ErrEmptyPassword
	}

	exists, err := i.repo.GetByEmail(ctx, email)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.String("email", email))
		return nil, er.ErrGetUserByEmail
	}
	if exists == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.String("email", email))
		return nil, er.ErrUserNotExists
	}

	err = bcrypt.CompareHashAndPassword([]byte(exists.Password), []byte(password))
	if err != nil {
		i.logger.Error(er.ErrIncorrectPassword.Error(), zap.String("email", email))
		return nil, er.ErrIncorrectPassword
	}

	newToken, err := i.token.NewToken(exists.ID, exists.Email)
	if err != nil {
		i.logger.Error(er.ErrNewToken.Error(), zap.Int64("id", exists.ID), zap.String("email", email))
		return nil, er.ErrNewToken
	}
	exists.Token = newToken

	ret, err := i.repo.UpdateToken(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrUpdateToken.Error(), zap.Int64("id", exists.ID), zap.String("email", email), zap.String("token", newToken))
		return nil, er.ErrUpdateToken
	}

	return ret, nil
}

func (i *impl) Logout(ctx contextx.Contextx) error {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return er.ErrUserNotExists
	}

	profile.Token = ""
	_, err := i.repo.UpdateToken(ctx, profile)
	if err != nil {
		i.logger.Error(er.ErrUpdateToken.Error(), zap.Error(err))
		return er.ErrUpdateToken
	}

	return nil
}
