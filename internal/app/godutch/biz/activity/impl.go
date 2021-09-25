package activity

import (
	"time"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "ActivityBiz")),
		repo:   repo,
		node:   node,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *event.Activity, err error) {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	ret, err := i.repo.GetByID(ctx, id, profile.ID)
	if err != nil {
		i.logger.Error(er.ErrGetActivityByID.Error(), zap.Error(err), zap.Any("user", profile), zap.Int64("act_id", id))
		return nil, er.ErrGetActivityByID
	}
	if ret == nil {
		i.logger.Error(er.ErrActivityNotExists.Error(), zap.Error(err), zap.Any("user", profile), zap.Int64("act_id", id))
		return nil, er.ErrActivityNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (infos []*event.Activity, total int, err error) {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, 0, er.ErrUserNotExists
	}

	if page <= 0 {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, 0, er.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, 0, er.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, profile.ID, size, (page-1)*size)
	if err != nil {
		i.logger.Error(er.ErrListActivities.Error(), zap.Error(err), zap.Any("user", profile), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrListActivities
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrActivityNotExists.Error(), zap.Any("user", profile), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrActivityNotExists
	}

	total, err = i.repo.Count(ctx, profile.ID)
	if err != nil {
		i.logger.Error(er.ErrCountActivity.Error(), zap.Error(err), zap.Any("user", profile))
		return nil, 0, er.ErrCountActivity
	}

	return ret, total, nil
}

func (i *impl) NewWithMembers(ctx contextx.Contextx, name string, emails []string) (info *event.Activity, err error) {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	if len(emails) == 0 {
		i.logger.Error(er.ErrEmptyEmail.Error())
		return nil, er.ErrEmptyEmail
	}

	if len(name) == 0 {
		i.logger.Error(er.ErrEmptyName.Error())
		return nil, er.ErrEmptyName
	}

	members, err := i.repo.GetByEmails(ctx, emails)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.Error(err), zap.Strings("emails", emails))
		return nil, er.ErrGetUserByEmail
	}

	created := &event.Activity{
		ID:      i.node.Generate().Int64() / 1000 * 1000,
		Name:    name,
		OwnerID: profile.ID,
		Owner: &user.Profile{
			ID:        profile.ID,
			Email:     profile.Email,
			Name:      profile.Name,
			CreatedAt: profile.CreatedAt,
		},
		Members:   members,
		CreatedAt: time.Now().Unix(),
	}
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateActivity.Error(), zap.Error(err), zap.Any("user", profile), zap.String("name", name), zap.Strings("emails", emails))
		return nil, er.ErrCreateActivity
	}

	return ret, nil
}

func (i *impl) InviteMembers(ctx contextx.Contextx, id int64, emails []string) (info *event.Activity, err error) {
	users, err := i.repo.GetByEmails(ctx, emails)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.Error(err), zap.Strings("emails", emails))
		return nil, er.ErrGetUserByEmail
	}

	ret, err := i.repo.AddMembers(ctx, id, users)
	if err != nil {
		i.logger.Error(er.ErrInviteMembers.Error(), zap.Error(err), zap.Any("users", users))
		return nil, er.ErrInviteMembers
	}

	return ret, nil
}

func (i *impl) ChangeName(ctx contextx.Contextx, id int64, name string) (info *event.Activity, err error) {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	if len(name) == 0 {
		i.logger.Error(er.ErrEmptyName.Error())
		return nil, er.ErrEmptyName
	}

	exists, err := i.repo.GetByID(ctx, id, profile.ID)
	if err != nil {
		i.logger.Error(er.ErrGetActivityByID.Error(), zap.Error(err), zap.Any("user", profile), zap.Int64("act_id", id))
		return nil, er.ErrGetActivityByID
	}
	if exists == nil {
		i.logger.Error(er.ErrActivityNotExists.Error(), zap.Any("user", profile), zap.Int64("act_id", id))
		return nil, er.ErrActivityNotExists
	}

	exists.Name = name
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrUpdateActivity.Error(), zap.Error(err), zap.Any("user", profile), zap.Any("exists", exists))
		return nil, er.ErrUpdateActivity
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return er.ErrUserNotExists
	}

	err := i.repo.Delete(ctx, id, profile.ID)
	if err != nil {
		i.logger.Error(er.ErrDeleteActivity.Error(), zap.Error(err), zap.Any("user", profile), zap.Int64("act_id", id))
		return er.ErrDeleteActivity
	}

	return nil
}
