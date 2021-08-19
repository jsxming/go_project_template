package service

import (
	"context"
	"go_project_template/global"
	"go_project_template/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DB)
	return svc
}
