package model

import (
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"go.uber.org/zap"
)

func (*UserDomain) DeleteUser(string) *rest_err.RestErr{
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
