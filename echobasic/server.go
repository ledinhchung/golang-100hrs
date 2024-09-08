package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"Name" form:"name"`
	Email string `json:"Email" form:"email"`
}

func hello(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Welcome "+name+" to Go world!\n")
}

func listUser(c echo.Context) error {
	return c.String(http.StatusOK, "List of user\n")
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Get user with id: "+id+"\n")
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Updated user id:"+id+"\n")
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Deleted user id: "+id+"\n")
}

func main() {
	e := echo.New()
	e.GET("/hello", hello)

	// User CURD
	e.GET("/users", listUser)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
