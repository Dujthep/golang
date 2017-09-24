package main

import (
	"strconv"
	"net/http"
	"public"
	"fizz_buzz"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/caesar/:txt", caesar)
	e.GET("/dec_caesar/:txt", deCaesar)
	e.GET("/letter_count/:txt", letterCount)
	e.GET("/fizz_buzz/:txt", fizzBuzz)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func caesar (c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string { "encrypted": public.Caesar(c.Param("txt"),4) })
}

func deCaesar (c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string { "decrypted": public.DecCaesar(c.Param("txt"),4) })
}

func letterCount (c echo.Context) error {
	return c.JSON(http.StatusOK, public.LettersCount(c.Param("txt")))
}

func fizzBuzz (c echo.Context) error {
	n, _ := strconv.Atoi(c.Param("txt"))
	return c.JSON(http.StatusOK, map[string]string { "fizz_buzz": fizz_buzz.CallFizzBuzz(n) } )
}