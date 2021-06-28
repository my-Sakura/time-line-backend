package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	timeLine "github.com/my-Sakura/time-line-backend/pkg/timeline/controller"
)

const (
	timeLineRouterGroup = "/api/v1/timeLine"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:123456@tcp(192.168.0.252:9090)/project?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	timeLineConn := timeLine.New(dbConn)

	timeLineConn.RegistRouter(router.Group(timeLineRouterGroup))

	log.Fatal(router.Run("0.0.0.0:10001"))
}
