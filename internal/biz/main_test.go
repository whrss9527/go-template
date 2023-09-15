package biz

import (
	"fmt"
	"os"
	"testing"

	"go-template/internal/data/model"
	fakeDb "go-template/internal/fake_db"
	"go-template/internal/pkg/db"
)

func setup() {
	fakeDb.InitFakeDb()
	fakeDb.InitTestRedis()

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
