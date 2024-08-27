package routers

import (
	"Back/Delivery/controllers"
	"Back/Infrastructure/auth_middleware"
	"Back/Repositories"
	usecases "blogapp/UseCase"
)

// refreshRouter
func RefreshTokenRouter() {
	refreshRouter := Router.Group("/refresh")
	{
		// generate new auth repo
		refreshrepo := Repositories.NewRefreshRepository(BlogCollections.RefreshTokens)
		refreshusecase := usecases.NewRefreshUseCase(refreshrepo)
		refreshcontroller := controllers.NewRefreshController(refreshusecase)

		refreshRouter.GET("", auth_middleware.AuthMiddleware(), refreshcontroller.Refresh)
	}
}
