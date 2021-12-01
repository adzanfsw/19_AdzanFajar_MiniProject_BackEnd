package route

import (
	"justrun/config"
	"justrun/controller"
	m "justrun/middleware"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteShoes() *echo.Echo {

	e := echo.New()

	auth := e.Group("auth")
	auth.POST("/token", controller.AuthLogin)

	jwtAuth := e.Group("jwt")
	jwtAuth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(config.JwtSecret),
	}))

	basicAuth := e.Group("basicAuth")
	basicAuth.Use(middleware.BasicAuth(m.BasicAuth))

	jwtAuth.POST("/api/shoes/add", controller.AddShoesController)
	jwtAuth.POST("/api/user/add", controller.AddUsersController)

	e.POST("/api/shoes-brand/add", controller.AddShoesBrandController)
	e.POST("/api/shoes-desc/add", controller.AddShoesDescController)
	e.POST("/api/wishlist/add", controller.AddWishController)
	e.POST("/api/reviews/add", controller.AddReviewController)

	jwtAuth.GET("/api/shoes", controller.GetShoesController)
	jwtAuth.GET("/api/shoes/:id", controller.ShoesbyIDController)

	jwtAuth.GET("/api/users", controller.GetUsersController)
	jwtAuth.GET("/api/users/:id", controller.UserbyIDController)

	e.GET("/api/shoes-brand", controller.GetShoesBrandController)
	e.GET("/api/shoes-type", controller.GetShoesTypeController)
	e.GET("/api/shoes-desc", controller.GetShoesDescController)

	e.GET("/api/reviews", controller.GetReviewsController)

	basicAuth.PUT("/api/shoes/update/:id", controller.UpdateShoesController)
	basicAuth.PUT("/api/users/update/:id", controller.UpdateUsersController)

	e.PUT("/api/reviews/update/:id", controller.UpdateReviewController)

	basicAuth.DELETE("/api/shoes/delete/:id", controller.DeleteShoesController)
	basicAuth.DELETE("/api/users/delete/:id", controller.DeleteUsersController)

	return e
}
