package test

import (
	"fmt"
	"go-template/internal/conf"
	"go-template/internal/data/model"
	"go-template/internal/pkg/db"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
)

var (
	dbName    = "mydb"
	tableName = "mytable"
	address   = "localhost"
	port      = 3380
)

func InitFakeDb() {
	go func() {
		Start()
	}()
	db.InitDbConfig(&conf.Data{
		Database: &conf.Data_Database{
			Dsn:             "no_user:@tcp(localhost:3380)/mydb?timeout=2s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4",
			ShowLog:         true,
			MaxIdleConn:     10,
			MaxOpenConn:     60,
			ConnMaxLifeTime: 4000,
		},
	})
	migrateTable()
}

func Start() {
	engine := sqle.NewDefault(
		memory.NewMemoryDBProvider(
			createTestDatabase(),
			information_schema.NewInformationSchemaDatabase(),
		))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	if err = s.Start(); err != nil {
		panic(err)
	}

}

func createTestDatabase() *memory.Database {
	db := memory.NewDatabase(dbName)
	db.EnablePrimaryKeyIndexes()
	return db
}

func migrateTable() {

	err := db.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

}
