package biz

import (
	"fmt"
	"go-template/internal/data/model"
	fake_mysql "go-template/internal/fake-mysql"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/redis"
	"os"
	"testing"
)

func setup() {
	fake_mysql.InitFakeDb()
	redis.InitTestRedis()

	result := db.DB.Create(&model.User{
		Id: "user1",
	})
	if result.Error != nil {
		panic(result.Error)
	}

	if result.Error != nil {
		panic(result.Error)
	}
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
