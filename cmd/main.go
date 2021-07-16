package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	account "github.com/my-Sakura/time-line-backend/pkg/account/controller"
	timeLine "github.com/my-Sakura/time-line-backend/pkg/timeline/controller"
)

const (
	timeLineRouterGroup = "/api/v1/timeLine"
	accountRouterGroup  = "/api/v1/account"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:123456@tcp(172.30.252.153:9092)/mysql?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	timeLineConn := timeLine.New(dbConn)
	accountConn := account.New(dbConn)

	timeLineConn.RegistRouter(router.Group(timeLineRouterGroup))
	accountConn.RegistRouter(router.Group(accountRouterGroup))

	log.Fatal(router.Run("0.0.0.0:10002"))
}
