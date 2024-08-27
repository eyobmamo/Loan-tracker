// router/router.go

package router

import (
    "github.com/gin-gonic/gin"
    "loan-tracker-api/controller"
    "loan-tracker-api/middleware"
)

func SetupRouter(userController *controller.UserController, adminController *controller.AdminController, loanController *controller.LoanController) *gin.Engine {
    router := gin.Default()

    // Public routes
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", userController.RegisterUser)
        userRoutes.GET("/verify-email", userController.VerifyEmail)
        userRoutes.POST("/login", userController.Login)
        userRoutes.POST("/token/refresh", userController.RefreshToken)
        userRoutes.POST("/password-reset", userController.RequestPasswordReset)
        userRoutes.POST("/password-reset", userController.UpdatePassword)
    }

    // Protected routes
    protectedRoutes := router.Group("/")
    protectedRoutes.Use(middleware.AuthMiddleware())
    {
        protectedRoutes.GET("/users/profile", userController.GetProfile)
        protectedRoutes.POST("/loans", loanController.ApplyForLoan)
        protectedRoutes.GET("/loans/:id", loanController.ViewLoanStatus)
    }

    // Admin routes
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        adminRoutes.GET("/users", adminController.GetAllUsers)
        adminRoutes.DELETE("/users/:id", adminController.DeleteUser)
        adminRoutes.GET("/loans", loanController.ViewAllLoans)
        adminRoutes.PATCH("/loans/:id/status", loanController.UpdateLoanStatus)
        adminRoutes.DELETE("/loans/:id", loanController.DeleteLoan)
    }

    return router
}
