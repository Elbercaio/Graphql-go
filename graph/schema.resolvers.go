package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/Elbercaio/gqlgen-todos/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := r.TodoDb.Create(input.Text, input.UserID)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserDb.Create(input.Name)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.TodoDb.List()
	if err != nil {
		return nil, err
	}
	var todosModel []*model.Todo
	for _, todo := range todos {
		todosModel = append(todosModel, &model.Todo{
			ID: todo.ID, Text: todo.Text, Done: todo.Done, // UserID: todo.User,
		})
	}
	return todosModel, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserDb.List()
	if err != nil {
		return nil, err
	}
	var usersModel []*model.User
	for _, user := range users {
		usersModel = append(usersModel, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return usersModel, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user, err := r.UserDb.GetByTodoId(obj.ID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	todos, err := r.TodoDb.GetByUserId(obj.ID)
	if err != nil {
		return nil, err
	}
	var todosModel []*model.Todo
	for _, todo := range todos {
		todosModel = append(todosModel, &model.Todo{
			ID: todo.ID, Text: todo.Text, Done: todo.Done, // UserID: todo.User,
		})
	}
	return todosModel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
