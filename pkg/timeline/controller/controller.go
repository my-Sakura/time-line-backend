package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// timeLineModel "github.com/time-line-backend/pkg/model/mysql"
)

var (
	deleted int8

	errDeleted         = errors.New("the user has been deleted")
	errActive          = errors.New("the user is not active")
	errUserIDNotExists = errors.New("Get Admin ID is not exists")
	errUserIDNotValid  = func(value interface{}) error {
		return fmt.Errorf("Get Admin ID is not valid. Is %s", value)
	}
)

type UserController struct {
	db  *sql.DB
	JWT *jwt.GinJWTMiddleware
}

func New(db *sql.DB) *UserController {
	uc := &UserController{
		db: db,
	}

	var err error

	uc.JWT, err = uc.newJWTMiddleware()
	if err != nil {
		log.Fatal(err)
	}

	return uc
}

func (uc *UserController) RegistRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter] server is nil")
	}

	err := mysql.CreateDatabase(uc.db)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/get", uc.get)
}

func (uc *UserController) get() {
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return 0, err
	}

	timeLine := SelectAllUnDeletedTimeLine(uc.db)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": timeLine})
}
