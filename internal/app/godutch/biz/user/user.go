package user

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare user's service function
type IBiz interface {
	// GetByID serve caller to given user id to get user information
	GetByID(ctx contextx.Contextx, id int64) (info *user.Profile, err error)

	// GetByToken serve caller to given user token to get user information
	GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error)

	// Signup serve caller to given email and password to register an account
	Signup(ctx contextx.Contextx, email, password, name string) (info *user.Profile, err error)

	// Login serve caller to given email and password to login system
	Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error)

	// Logout serve caller to logout the system
	Logout(ctx contextx.Contextx) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
