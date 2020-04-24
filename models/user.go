package models

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

// AddUsers add new user to user slice and return the new user or error
func AddUsers(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}