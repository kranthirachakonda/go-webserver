package models

// User represents a webservice user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers lists all users on the system storage
func GetUsers() []*User {
	return users
}

// AddUser adds a new user to the system storage
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
