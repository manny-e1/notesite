package user

import (
	"github.com/ns/models"
)


//UserService ... this is our service usescase(has (interface)abstract classes that outer layers can use)
type UserService interface {
	RegisterUser(user *models.User)(*models.User,[]error)
	Users() ([]models.User,[]error)
	GetUser(int uint) (*models.User,[]error)
    GetUserByUsername(username string) (*models.User, []error)
	EditUser(user *models.User)(*models.User,[]error)
	DeleteUser(id uint)(*models.User,[]error)
	UsernameExists(username string) bool
}


