package model

import (
	"github.com/jeanroths/CRUD-auth/src/controller"
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	) UserDomainInterface {
		return &UserDomain{
			email, password, name,
	}
}

type UserDomain struct {
	Name string 
	Email string 
	Password string 
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *rest_err.RestErr
	UpdateUser(string, UserDomain) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}