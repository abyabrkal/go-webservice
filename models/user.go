package models

import (
	"fmt"
	"errors"
)

// User structure 
type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

// GetUsers return the pointer slice of address to user structs
func GetUsers() []*User {
	return users
}

// AddUser add new user to user slice and return the new user or error
func AddUser(u User) (User, error) {
	if(u.ID != 0) {
		return User{}, errors.New("New user must not include id or it must be")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserById finds user by ID provided
func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' is not found", id)
}


// UpdateUser updates a User from provided user data
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' is not found", u.ID)
}


// RemoveUserById removes a user based on id
func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' is not found", id)
}