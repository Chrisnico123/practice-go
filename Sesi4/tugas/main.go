package main

import (
	"echo/handler"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)



func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceID := c.Get("trace_id")
		log.Printf("%s message=\"incoming request\" method=%s uri=%s trace_id=%s\n", time.Now().Format(time.RFC3339), c.Request().Method, c.Request().URL.Path, traceID)
		err := next(c)
		return err
	}
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceID := c.Get("trace_id")
		log.Printf("%s message=\"incoming request\" trace_id=%s\n", time.Now().Format(time.RFC3339), traceID)
		err := next(c)
		if err != nil {
			log.Printf("%s message=\"error\" method=%s uri=%s trace_id=%s\n", time.Now().Format(time.RFC3339), c.Request().Method, c.Request().URL.Path, traceID)
		}
		log.Printf("%s message=\"finish request\" method=%s uri=%s trace_id=%s\n", time.Now().Format(time.RFC3339), c.Request().Method, c.Request().URL.Path, traceID)
		return err
	}
}

func main() {
	router := echo.New()

	user := router.Group("users")
	user.Use(Middleware)
	user.Use(Logger)
	{
		user.GET("" , handler.GetDataAllUser)
		user.POST("" , handler.CreateDataUser)
		user.PUT("/:id", handler.UpdateDataUser)
		user.DELETE("/:id", handler.DeleteUser)
	}

	router.Start(":3000")
}