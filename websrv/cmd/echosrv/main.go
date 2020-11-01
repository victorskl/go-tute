package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func listUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getVersion(c echo.Context) error {
	versionMap := map[string]string{"version": "0.1.0"}
	return c.JSON(http.StatusOK, versionMap)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("/users/:id", getUser)

	e.GET("/users", listUsers)
	e.GET("/version", getVersion)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Using https://echo.labstack.com/
// More https://github.com/labstack/echox/tree/master/cookbook
