package routers

import (
	"Back/Delivery/controllers"
	"Back/Infrastructure/auth_middleware"
	"Back/Repositories"
	
)

func AuthRouter() {
	authRouter := Router.Group("/users")
	{
		userRepo := Repositories.NewUserRepository(LoanCollections.Users, LoanCollections.RefreshTokens)

		// generate new auth repo
		authrepo := Repositories.NewAuthRepository(LoanCollections.Users, LoanCollections.RefreshTokens, userRepo)
		authusecase := usecases.NewAuthUseCase(authrepo)
		authcontroller := controllers.NewAuthController(authusecase)

		// register
		authRouter.POST("/register", authcontroller.Register)
		//login
		authRouter.POST("/login", authcontroller.Login)

		// oauth login with google
		authRouter.GET("/login/google", authcontroller.LoginHandlerGoogle)
		authRouter.GET("/callback", authcontroller.CallbackHandler)

		//logout
		authRouter.GET("/logout", auth_middleware.AuthMiddleware(), authcontroller.Logout)
		// forget password
		authRouter.POST("/password-reset", authcontroller.ForgetPassword)
		// authRouter.GET("/forget-password/:reset_token", authcontroller.ForgetPasswordForm)
		authRouter.POST("/password-reset/:reset_token", authcontroller.ResetPassword)

		// activate account
		authRouter.GET("/activate/:activation_token", authcontroller.ActivateAccount)

	}
}
