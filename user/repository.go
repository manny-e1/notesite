package user

import (
	"github.com/ns/models"
)


//UserRepository repository(interface) specifies User user related database operations
type UserRepository interface {
    Users() ([]models.User,[]error)
    GetUser(id uint) (*models.User, []error)
    GetUserByUsername(username string) (*models.User,[]error)
	RegisterUser(user *models.User)(*models.User,[]error)
	EditUser(user *models.User)(*models.User,[]error)
	DeleteUser(id uint)(*models.User,[]error)
	UsernameExists(username string) bool

}


