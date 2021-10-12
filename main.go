package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Memo string
}

func main() {
	r := gin.Default()
	r.Static("styles", "./styles")
	r.LoadHTMLGlob("templates/*")
	dbInit()
	r.GET("/", getHandler)
	r.POST("/new", postHandler)
	r.Run()
}

func getHandler(c *gin.Context) {
	todo := getAll()
	c.HTML(200, "index.html", gin.H{"todo": todo})
}

func postHandler(c *gin.Context) {
	memo := c.PostForm("memo")
	create(memo)
	c.Redirect(302, "/")
}

// db migrate
func dbInit() {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("faild to connect database \n")
	}
	db.AutoMigrate(&Todo{})
}

// create database

func create(memo string) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("fail to connect database \n")
	}
	db.Create(&Todo{Memo: memo})
}

func getAll() []Todo {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("fail to connect database \n")
	}
	var todo []Todo
	db.Find(&todo)
	return todo
}
