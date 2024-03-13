package models

import (
	"github.com/jeanroths/CRUD-auth/src/configuration/rest_err"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"go.uber.org/zap"
)

func (*UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr){
	logger.Info("Init findUser model", zap.String("journey", "findUser"))
	return nil
}