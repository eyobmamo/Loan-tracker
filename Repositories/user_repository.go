package Repositories

import (
	ps "Back/Infrastructure/password_services"
	"sync"

	// "time"

	"Back/Domain"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator"
)

type userRepository struct {
	validator       *validator.Validate
	collection      Domain.Collection
	TokenRepository Domain.RefreshRepository
	mu              sync.RWMutex
}

func NewUserRepository(_collection Domain.Collection, token_collection Domain.Collection) *userRepository {
	return &userRepository{
		validator:       validator.New(),
		collection:      _collection,
		TokenRepository: NewRefreshRepository(token_collection),
		mu:              sync.RWMutex{},
	}
}

// create user
func (us *userRepository) CreateUser(ctx context.Context, user *Domain.User) (Domain.OmitedUser, error, int) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	// Check if user email is taken
	existingUserFilter := bson.D{{"email", user.Email}}
	existingUserCount, err := us.collection.CountDocuments(ctx, existingUserFilter)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	if existingUserCount > 0 {
		return Domain.OmitedUser{}, errors.New("Email is already taken"), http.StatusBadRequest
	}
	// User registration logic
	hashedPassword, err := ps.GenerateFromPasswordCustom(user.Password)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	user.Password = string(hashedPassword)
	insertResult, err := us.collection.InsertOne(ctx, user)
	if err != nil {
		return Domain.OmitedUser{}, err, 500
	}
	// Fetch the inserted task
	var fetched Domain.OmitedUser
	err = us.collection.FindOne(context.TODO(), bson.D{{"_id", insertResult.InsertedID.(primitive.ObjectID)}}).Decode(&fetched)
	if err != nil {
		fmt.Println(err)
		return Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	if fetched.Email != user.Email {
		return Domain.OmitedUser{}, errors.New("User Not Created"), 500
	}
	fetched.Password = ""
	return fetched, nil, 200
}

// get all users
func (us *userRepository) GetUsers(ctx context.Context) ([]*Domain.OmitedUser, error, int) {
	us.mu.RLock()
	defer us.mu.RUnlock()
	var results []*Domain.OmitedUser

	// Pass these options to the Find method
	findOptions := options.Find()
	// findOptions.SetLimit(2)
	filter := bson.D{{}}

	// Here's an array in which you can store the decoded documents

	// Passing bson.D{{}} us the filter matches all documents in the collection
	cur, err := us.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal("error in finding users", err)
		log.Fatal(err)
		return []*Domain.OmitedUser{}, err, 0
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem Domain.OmitedUser
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error in decoding user", err)
			fmt.Println(err.Error())
			// #handelthislater
			// should this say there was a decoding error and return?
			return []*Domain.OmitedUser{}, err, 500
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return []*Domain.OmitedUser{}, err, 500
	}

	// Close the cursor once finished
	cur.Close(ctx)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil, 200
}

// get user by id
func (us *userRepository) GetUsersById(ctx context.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	var filter bson.D
	filter = bson.D{{"_id", id}}
	var result Domain.OmitedUser
	err := us.collection.FindOne(ctx, filter).Decode(&result)
	// # handel this later
	if err != nil {
		return Domain.OmitedUser{}, errors.New("User not found"), http.StatusNotFound
	}
	if current_user.Role == "user" && result.ID != current_user.ID {
		return Domain.OmitedUser{}, errors.New("permission denied"), http.StatusForbidden

	}
	return result, nil, 200
}
