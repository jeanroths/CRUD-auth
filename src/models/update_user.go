package models

import (
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"go.uber.org/zap"
)

func (*UserDomain) UpdateUser(string, UserDomain) *rest_err.RestErr{
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))
	return nil
}