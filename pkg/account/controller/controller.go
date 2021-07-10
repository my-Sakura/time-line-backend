package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	accounts = map[string]string{
		"root":       "123456",
		"sakura":     "123456",
		"zhangbokai": "123456",
		"wangzequan": "123456",
		"chaijiayin": "123456",
	}
)

type AccountController struct {
	db *sql.DB
}

func New(db *sql.DB) *AccountController {
	ac := &AccountController{
		db: db,
	}

	return ac
}

func (ac *AccountController) RegistRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter] server is nil")
	}

	r.POST("/login", ac.login)
}

func (ac *AccountController) login(c *gin.Context) {
	var req struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest})
		return
	}

	if password, ok := accounts[req.UserName]; ok {
		if req.Password == password {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
			return
		}
	}

	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "data": "login failed"})
}
