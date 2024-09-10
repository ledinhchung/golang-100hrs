package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dbname := "basic.db"

	db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func hello(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Welcome "+name+" to Go world!\n")
}

func getUsers(c echo.Context) error {
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(User)
	var oldUser User
	if err := c.Bind(user); err != nil {
		return err
	}
	db.First(&oldUser, id)
	db.Model(&oldUser).Update("Name", user.Name)
	db.Model(&oldUser).Update("Email", user.Email)
	return c.JSON(http.StatusOK, oldUser)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	db.Delete(&user, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	e.GET("/hello", hello)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User CURD
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
