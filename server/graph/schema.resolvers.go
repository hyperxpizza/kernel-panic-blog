package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/goware/emailx"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/graph/generated"
	"github.com/hyperxpizza/kernel-panic-blog/server/graph/model"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	usernameExists := database.CheckIfUsernameExists(username)
	if usernameExists == false {
		return nil, fmt.Errorf("Username does not exist")
	}

	//compare passwords
	passwordToCheck := database.GetUsersPassword(username)
	if passwordToCheck == "" {
		return nil, fmt.Errorf("Error while getting password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordToCheck), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Passwords do not match")
	}

	id := database.GetUsersID(username)

	//Generate token
	token, err := middleware.GetAuthToken(id)
	if err != nil {
		return nil, fmt.Errorf("Error while generating token")
	}

	payload := model.AuthPayload{
		Token:  token,
		UserID: id.String(),
	}

	return &payload, nil
}

func (r *mutationResolver) InsertUser(ctx context.Context, username string, password1 string, password2 string, email string) (*model.AuthPayload, error) {
	//validate email
	err := emailx.Validate(email)
	if err != nil {
		if err == emailx.ErrInvalidFormat {
			return nil, fmt.Errorf("Invalid email format")
		}

		return nil, fmt.Errorf("Error while validating email")
	}

	//TODO: Add username validation

	//Check if username taken
	usernameExists := database.CheckIfUsernameExists(username)
	if usernameExists == true {
		return nil, fmt.Errorf("Username already taken")
	}

	emailExists := database.CheckIfEmailTaken(email)
	if emailExists == true {
		return nil, fmt.Errorf("Email already taken")
	}

	// Compare passwords
	if password1 != password2 {
		fmt.Errorf("Passwords do not match")
	}

	//Create PasswordHash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password1), 10)
	if err != nil {
		return nil, fmt.Errorf("Error while hashing password")
	}

	err = database.InsertUser(username, string(passwordHash), email)
	if err != nil {
		return nil, err
	}

	// Get ID
	id := database.GetUsersID(username)

	//Generate token
	token, err := middleware.GetAuthToken(id)
	if err != nil {
		return nil, fmt.Errorf("Error while generating token")

	}

	payload := model.AuthPayload{
		Token:  token,
		UserID: id.String(),
	}

	return &payload, nil
}

func (r *mutationResolver) InsertPost(ctx context.Context, title string, subtitle *string, content string, authorID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	//posts, err := database.GetAllPosts()
	//if err != nil {
	//	return nil, fmt.Errorf("error")
	//}

	//return &posts, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
