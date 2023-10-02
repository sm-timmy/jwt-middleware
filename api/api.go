package api

import (
	"github.com/labstack/echo/v4"
	"local/controller"
	"local/middleware"
)

// setupRouter sets up the router and adds the routes.
func SetupRouter() *echo.Echo {
	// Create a new router
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome To This Website")
	})
	api := e.Group("/api")
	public := api.Group("/public")
	public.POST("/login", controller.Login)
	public.POST("/signup", controller.Signup)

	protected := api.Group("/protected")
	protected.Use(middleware.Authz)
	protected.GET("/users", controller.GetUsers)
	protected.POST("/users", controller.SaveUser)
	protected.GET("/users/:id", controller.GetUser)
	protected.PUT("/users", controller.UpdateUser)
	protected.DELETE("/users/:id", controller.DeleteUser)
	return e
}
