package router

import (
	interfaces2 "Spy-Cat-Agency/src/internal/missions/interfaces"
	handlers2 "Spy-Cat-Agency/src/internal/missions/interfaces/handlers"
	"Spy-Cat-Agency/src/internal/shared/middlewares"
	"Spy-Cat-Agency/src/internal/spycats/interfaces"
	"Spy-Cat-Agency/src/internal/spycats/interfaces/handlers"

	_ "Spy-Cat-Agency/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(spycatsHandler *handlers.SpyCatHandler, missionHandler *handlers2.MissionHandler) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middlewares.LoggingMiddleware())

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-length"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	interfaces.SetUpSpyCatRouter(api, spycatsHandler)
	interfaces2.SetUpMissionRouter(api, missionHandler)

	return r

}
