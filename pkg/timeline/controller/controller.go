package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/my-Sakura/time-line-backend/pkg/timeline/model/mysql"
)

type TimeLineController struct {
	db *sql.DB
}

func New(db *sql.DB) *TimeLineController {
	tc := &TimeLineController{
		db: db,
	}

	return tc
}

func (tc *TimeLineController) RegistRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter] server is nil")
	}

	err := mysql.CreateDatabase(tc.db)
	if err != nil {
		log.Fatal(err)
	}

	err = mysql.CreateTimeLine(tc.db)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/get", tc.get)
	r.GET("/getByLabel", tc.getByLabel)

	r.POST("/add", tc.add)
	r.POST("/delete", tc.delete)
	r.POST("/update", tc.update)
}

func (tc *TimeLineController) get(c *gin.Context) {
	timeLine, err := mysql.SelectAllUnDeletedTimeLine(tc.db)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"error": http.StatusInternalServerError, "data": timeLine})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": timeLine})
}

func (tc *TimeLineController) getByLabel(c *gin.Context) {
	var req struct {
		Label string `json:"label"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	timeLine, err := mysql.SelectByColorUnDeletedTimeLine(tc.db, req.Label)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"error": http.StatusInternalServerError, "data": timeLine})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": timeLine})
}

func (tc *TimeLineController) add(c *gin.Context) {
	var req struct {
		Title     string    `json:"title"`
		Value     string    `json:"value"`
		Label     string    `json:"label"`
		EventTime time.Time `json:"event_time"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	err = mysql.InsertTimeLine(tc.db, req.Title, req.Value, req.Label, req.EventTime)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (tc *TimeLineController) delete(c *gin.Context) {
	var req struct {
		ID uint32 `json:"id"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError})
		return
	}

	err = mysql.DeleteTimeLine(tc.db, req.ID)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (tc *TimeLineController) update(c *gin.Context) {
	var req struct {
		ID        uint32    `json:"id"`
		Title     string    `json:"title"`
		Value     string    `json:"value"`
		Label     string    `json:"label"`
		Color     string    `json:"color"`
		EventTime time.Time `json:"event_time"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError})
		return
	}
	fmt.Println(req, "--")
	err = mysql.UpdateTimeLineByID(tc.db, req.ID, req.Title, req.Value, req.Label, req.Color, req.EventTime)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
