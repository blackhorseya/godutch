package activity

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	idGen  *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, idGen *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "ActivityBiz")),
		repo:   repo,
		idGen:  idGen,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *event.Activity, err error) {
	// todo: 2021-09-24|11:48|Sean|impl me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (infos []*event.Activity, total int, err error) {
	// todo: 2021-09-24|11:48|Sean|impl me
	panic("implement me")
}

func (i *impl) NewWithMembers(ctx contextx.Contextx, name string, email []string) (info *event.Activity, err error) {
	// todo: 2021-09-24|11:48|Sean|impl me
	panic("implement me")
}

func (i *impl) ChangeName(ctx contextx.Contextx, id int64, name string) (info *event.Activity, err error) {
	// todo: 2021-09-24|11:48|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	// todo: 2021-09-24|11:48|Sean|impl me
	panic("implement me")
}
