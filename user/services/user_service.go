package services

import (
	"github.com/ns/models"
	"github.com/ns/user"
)

//UserServiceImpl implements user.UserService interface
type UserServiceImpl struct {
	userRepo user.UserRepository
}

//NewUserServiceImpl ... creates an object of UserServiceImpl
func NewUserServiceImpl(UserRep user.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: UserRep}
}


func (usi *UserServiceImpl) Users() ([]models.User, []error) {
	usrs, err := usi.userRepo.Users()
	if len(err) > 0 {
		return nil, err
	}
	return usrs, err
}

//GetUser ... returns one user row with the given username
func (usi *UserServiceImpl) GetUser(id uint) (*models.User, []error) {
	//check username?
	user, err := usi.userRepo.GetUser(id)
	if len(err) > 0 {
		return nil, err
	}
	return user, nil
}

func (usi *UserServiceImpl) GetUserByUsername(username string) (*models.User, []error){
   user, err := usi.userRepo.GetUserByUsername(username)
   	if len(err) > 0 {
   		return nil, err
   	}
   	return user, err
}


//RegisterUser ... registers a new user
func (usi *UserServiceImpl) RegisterUser(user *models.User) (*models.User,[]error) {
	user,err := usi.userRepo.RegisterUser(user)
	if len(err) > 0 {
		return user,err
	}
	return user,nil
}
//EditUser ... edit existing user data(profile)
func (usi *UserServiceImpl) EditUser(users *models.User)(*models.User,[]error) {
	user,err := usi.userRepo.EditUser(users)
	if len(err) > 0 {
		return user,err
	}
	return user,nil
}

//DeleteUser ... delete existing user with the given id
func (usi *UserServiceImpl) DeleteUser(id uint)(*models.User,[]error) {
	user,err := usi.userRepo.DeleteUser(id)
	if len(err) > 0 {
		return nil,err
	}
	return user,nil
}

func (usi *UserServiceImpl) UsernameExists(username string) bool {
	exists := usi.userRepo.UsernameExists(username)
	return exists
}

