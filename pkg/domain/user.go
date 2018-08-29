package domain

import "github.com/tokopedia/tdk/go/app/resource"

type UserDomain struct {
	resource UserResourceItf
}

func InitUserDomain(rsc UserResourceItf) UserDomain {
	return UserDomain{
		resource: rsc,
	}
}

func (user UserDomain) IsValidUser(userID int) bool {
	if err := user.resource.FindUser(userID); err != nil {
		return false
	}
	return true
}

type UserResourceItf interface {
	FindUser(int) error
}

type UserResource struct {
	DB resource.SQLDB
}

// this is just simple function to give a picture of how user resource works
func (user UserResource) FindUser(userId int) error {
	// you may check to db is this userId present in database
	return nil
}
