package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

// User ...
type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func userHandler(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func userformHandler(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}

func userqueryHandler(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, name)
}

func userNameHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Request!")
	})
	e.POST("/users", userHandler)
	e.POST("/userform", userformHandler)
	e.GET("/userquery", userqueryHandler)
	e.GET("/users/:name", userNameHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
