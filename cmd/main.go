package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	timeLine "github.com/my-Sakura/time-line-backend/pkg/timeline/controller"
)

const (
	timeLineRouterGroup = "/api/v1/timeLine"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:123456@tcp(123.56.162.98:9092)/mysql?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	timeLineConn := timeLine.New(dbConn)

	timeLineConn.RegistRouter(router.Group(timeLineRouterGroup))

	log.Fatal(router.Run("0.0.0.0:10002"))
}

docker volume create timeline-mysql-config
docker volume create timeline-mysql-data
docker volume create timeline-mysql-log
docker run --name timeline -p 9092:3306 \ 
  --mount type=volume,src=timeline-mysql-config,dst=/etc/mysql \
  --mount type=volume,src=timeline-mysql-data,dst=/var/lib/mysql \
  --mount type=volume,src=timeline-mysql-log,dst=/var/log/mysql \
  -e MYSQL\_ROOT\_PASSWORD=123456 \
	-d mysql:5.7
	