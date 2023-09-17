package routes

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// Returns a new router
func NewRouter() (*fizz.Fizz, error) {
	// gin instance
	engine := gin.Default()
	// default cors config, Allow Origin, Authorization header
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	engine.Use(cors.New(config))

	// html
	engine.LoadHTMLGlob("html/*.html")
	// index
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// error 404
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})

	// fizz instance
	fizz := fizz.NewFromEngine(engine)

	// OpenApi info
	infos := &openapi.Info{
		Title:       "Trek server",
		Description: "Best choice for tracking your motorcycle or whatever.",
		Version:     "1.0.0",
	}

	// base API route
	grp := fizz.Group("api", "", "")

	// OpenAPI spec
	grp.GET("openapi.json", nil, fizz.OpenAPI(infos, "json"))

	// TODO setup other routes

	// TODO login
	// TODO create user
	// TODO get user(s)
	// TODO delete user(s)
	// TODO add tracker and generate API key
	// TODO regenerate API key
	// TODO get tracker(s)
	// TODO delete tracker(s)

	// TODO GNSS data

	if len(fizz.Errors()) != 0 {
		return nil, fmt.Errorf("errors: %v", fizz.Errors())
	}
	return fizz, nil
}
