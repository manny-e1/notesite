package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/ns/models"
)

//UserRepositoryImpl ... implements the User.UserRepository interface
type UserRepositoryImpl struct {
	conn *gorm.DB
}

//NewUserRepositoryImpl will create an object of  UserRepositoryImpl
func NewUserRepositoryImpl(Conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}

//retrieve all the users
func (usrRep *UserRepositoryImpl) Users() ([]models.User, []error){
    usrs := []models.User{}
    err := usrRep.conn.Find(&usrs).GetErrors()
    if len(err) > 0 {
        return nil, err
    }
    return usrs, nil
}

//retrieve an user with a given ID
func (usrRep *UserRepositoryImpl) GetUser(id uint) (*models.User,[]error){
    user := models.User{}
    err := usrRep.conn.First(&user,id).GetErrors()
    if len(err) > 0{
        return nil,err    
    }
    return &user,nil
}

//retrieve an user with a give username
func (usrRep *UserRepositoryImpl) GetUserByUsername(username string) (*models.User,[]error){
    user := models.User{}
    err := usrRep.conn.Find(&user,"username=?",username).GetErrors()
    if len(err) > 0{
        return nil,err
    }
    return &user,nil
}

//registers a new user to a db
func (usrRep *UserRepositoryImpl) RegisterUser(usr *models.User) (*models.User, []error){
    user := usr
    err := usrRep.conn.Create(user).GetErrors()
    if len(err) > 0{
        return nil,err
    }
    return user,nil
}

//Edits the user's info
func (usrRep *UserRepositoryImpl) EditUser(usr *models.User) (*models.User,[]error){
    user := usr
    err:=usrRep.conn.Debug().Save(user).GetErrors()
    //err := usrRep.conn.Raw("update users set username, name, email, about, password, image where id = ?",id ).Scan(&user).GetErrors()
    if len(err)>0{
        return nil,err
    }
    return user,nil
}

//deletes a given user
func (usrRep *UserRepositoryImpl) DeleteUser(id uint) (*models.User,[]error){
    user,err := usrRep.GetUser(id)
    if len(err)>0{
        return nil,err
    }
    err = usrRep.conn.Delete(user,id).GetErrors()
    if len(err)>0{
        return nil,err
    }
    return user,nil
}

//checks if there's already a user by that username
func (usrRep *UserRepositoryImpl) UsernameExists(username string) bool{
    user := models.User{}
    	errs := usrRep.conn.Find(&user, "username=?", username).GetErrors()
    	if len(errs) > 0 {
    		return false
    	}
    	return true
}
