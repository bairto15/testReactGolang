package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

type Table struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Struct []string `json:"struct"`
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(LiberalCORS)
	router.GET("/", h.hello)
	router.GET("/table", h.getTable)
	router.POST("/table", h.setTable)

	return router
}

func (h *Handler) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1.001",
	})
}

func (h *Handler) getTable(c *gin.Context) {
	const path = "tables.json"
	data := make([]Table, 0)

	read, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal(read, &data)

	c.JSON(http.StatusOK, data)
}

//При одновременном редактировании сохраниться последнее сохранение
func (h *Handler) setTable(c *gin.Context) {
	var data []Table

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const path = "tables.json"
	file, _ := json.MarshalIndent(data, "", " ")
	ioutil.WriteFile(path, file, 0644)

	c.JSON(http.StatusOK, gin.H{})
}

//Отключение cors
func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers",
				c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
