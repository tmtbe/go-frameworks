package testcontainer

import (
	"context"
	"github.com/google/wire"
	"io/ioutil"
	"strings"
	app_context "test/internal/app/context"
)

func NewTSContext() context.Context {
	background := context.Background()
	return background
}

type Background struct {
	app_context.Context
	TestContainersContext context.Context
}

func (bg *Background) MustSetUpDb(name string) {
	sqls, err := ioutil.ReadFile("./sql/" + name + ".sql")
	if err != nil {
		panic(err)
	}
	bg.MustSetUpDbWithSql(string(sqls))
}

func (bg *Background) MustSetUpDbWithSql(stmt string) {
	stmt = strings.TrimSuffix(stmt, "\n")
	stmt = strings.TrimSuffix(stmt, " ")
	stmt = strings.TrimSuffix(stmt, ";")
	if _, err := bg.DB.Exec(stmt); err != nil {
		panic(err)
	}
}

var ProviderSet = wire.NewSet(NewTSContext, wire.Struct(new(Background), "*"))
