package models

import (
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	return nil
}