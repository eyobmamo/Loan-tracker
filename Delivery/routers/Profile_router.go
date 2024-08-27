package routers

import (
	"Back/Delivery/controllers"
	"Back/Infrastructure/auth_middleware"
	"Back/Repositories"
	usecases "Back/UseCase"
)

func ProfileRouter() {
	profileRouter := Router.Group("/user")
	profileRouter.Use(auth_middleware.AuthMiddleware())
	{

		// generate new auth repo
		profile_repo := Repositories.NewProfileRepository(LoanCollections.Users, LoanCollections.RefreshTokens)

		profile_usecase := usecases.NewProfileUseCase(profile_repo)
		profile_controller := controllers.NewProfileController(profile_usecase)

		// get all users
		profileRouter.GET("/", profile_controller.GetProfile)
		profileRouter.PATCH("/", profile_controller.UpdateProfile)
		profileRouter.DELETE("/", profile_controller.DeleteProfile)

	}
}
