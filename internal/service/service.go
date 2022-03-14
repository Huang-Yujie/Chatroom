package service

import (
	"context"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
