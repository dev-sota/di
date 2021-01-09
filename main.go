package main

import (
	"database/sql"

	"github.com/dev-sota/di/repository"
	"github.com/dev-sota/di/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)

	// example
	us.FindById(1)
}
