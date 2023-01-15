package api

import (
	"github.com/techotron/dog-ceo-api/backend/controllers"
	"github.com/techotron/dog-ceo-api/backend/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

// Setup API using gin with routes and cors settings
func Setup() *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(log.Logger), gin.Recovery())
	addAllowAllCors(router)
	addRoutes(router)
	return router
}

// addAllowAllCors sets cors config
func addAllowAllCors(g *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	g.Use(cors.New(corsConfig))
}

// addRoutes using existing controllers
func addRoutes(g *gin.Engine) {
	g.GET("/info", controllers.GetInfo)

	// breeds/list/all
	// breeds/image/random (random image of a dog)
	// breeds/image/random/:number (returns :number of random collection of images)
	// breed/:breed/images (list of all images for :breed)
	// breed/:breed/images/random (random image for :breed)
	// breed/:breed/images/random/:number (returns :number of random collection of images from :breed)
	// breed/:breed/list (returns list of sub-breeds for :breed)
	// breed/:breed/:subbreed/images (returns list of images for given :breed / :subbreed)
	// breed/:breed/:subbreed/images/random (random image for :breed / :subbreed)
	// breed/:breed/:subbreed/images/random/:number (returns :number of random collection of images from :breed / :subbreed)

}
