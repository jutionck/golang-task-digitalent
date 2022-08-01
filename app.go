package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-task-digitalent/config"
	"github.com/jutionck/golang-task-digitalent/model"
	"github.com/jutionck/golang-task-digitalent/repository"
	"github.com/jutionck/golang-task-digitalent/usecase"
)

func main() {

	// Panggil Koneksi
	db := config.NewConfig()
	repo := repository.NewTaskRepository(db.Db)
	useCase := usecase.NewTaskUseCase(repo)

	router := gin.Default()
	rg := router.Group("/api")

	rg.GET("/task", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		tasks, err := useCase.FindAll()
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			c.JSON(200, tasks)
		}
	})

	rg.POST("/task", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		err := useCase.RegisterNewTask(&task)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			c.JSON(200, task)
		}
	})

	rg.GET("/task/:id", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		id := c.Param("id")
		task, err := useCase.FindById(id)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			c.JSON(200, task)
		}
	})

	rg.PUT("/task", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		err := useCase.UpdateTask(&task)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			c.JSON(200, task)
		}
	})

	rg.DELETE("/task/:id", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		id := c.Param("id")
		err := useCase.DeleteTask(id)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			c.JSON(200, "deleted success")
		}
	})

	port := os.Getenv("API_PORT")
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
