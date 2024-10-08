package Domain

import (
	"Back/Dtos"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	Login(ctx context.Context, user *User) (Tokens, error, int)
	Register(ctx context.Context, user *Dtos.RegisterUserDto) (*OmitedUser, error, int)
	Logout(ctx context.Context, user_id primitive.ObjectID) (error, int)
	GoogleLogin(ctx context.Context) string
	CallbackHandler(ctx context.Context, code string) (Tokens, error, int)
	GenerateTokenFromUser(ctx context.Context, existingUser User) (Tokens, error, int)
	ResetPassword(ctx context.Context, email string, password string, resetToken string) (error, int)
	ForgetPassword(ctx context.Context, email string) (error, int)
	ActivateAccount(ctx context.Context, token string) (error, int)
	SendActivationEmail(email string) (error, int)
}

type AuthUseCase interface {
	Login(c *gin.Context, user *User) (Tokens, error, int)
	Register(c *gin.Context, user *Dtos.RegisterUserDto) (*OmitedUser, error, int)
	Logout(c *gin.Context, user_id primitive.ObjectID) (error, int)
	GoogleLogin(c *gin.Context) string
	CallbackHandler(c *gin.Context, code string) (Tokens, error, int)
	ResetPassword(c *gin.Context, email string, password string, resetToken string) (error, int)
	ForgetPassword(c *gin.Context, email string) (error, int)
	ActivateAccount(c *gin.Context, token string) (error, int)
}

type RefreshRepository interface {
	UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int)
	FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}
type RefreshUseCase interface {
	// UpdateToken(c *gin.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(c *gin.Context, userid primitive.ObjectID) (error, int)
	FindToken(c *gin.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(c *gin.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (OmitedUser, error, int)
	GetUsers(ctx context.Context) ([]*OmitedUser, error, int)
	GetUsersById(ctx context.Context, id primitive.ObjectID, user AccessClaims) (OmitedUser, error, int)
	UpdateUsersById(ctx context.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteUsersById(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
	PromoteUser(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	DemoteUser(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	ChangePassByEmail(ctx context.Context, email string, password string) (OmitedUser, error, int)
	FindByEmail(ctx context.Context, email string) (OmitedUser, error, int)
}

type UserUseCases interface {
	CreateUser(c *gin.Context, user *User) (OmitedUser, error, int)
	GetUsers(c *gin.Context) ([]*OmitedUser, error, int)
	GetUsersById(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	UpdateUsersById(c *gin.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteUsersById(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
	PromoteUser(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	DemoteUser(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
}

type ProfileUseCases interface {
	GetProfile(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (OmitedUser, error, int)
	UpdateProfile(c *gin.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteProfile(c *gin.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
}

type ProfileRepository interface {
	GetProfile(ctx context.Context, id primitive.ObjectID, user AccessClaims) (OmitedUser, error, int)
	UpdateProfile(ctx context.Context, id primitive.ObjectID, user User, current_user AccessClaims) (OmitedUser, error, int)
	DeleteProfile(ctx context.Context, id primitive.ObjectID, current_user AccessClaims) (error, int)
}
