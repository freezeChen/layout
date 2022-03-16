package data

import (
	"gitee.com/bethink1501/24on-library/db"
	"github.com/freezeChen/layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"xorm.io/xorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *xorm.Engine
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	initDb := db.InitDb(c.Mysql)

	cleanup := func() {
		initDb.Close()
		log.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
