package user

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func AddUser(ctx context.Context, req *library.AddUserReq) (*library.AddUserResp, error) {
	userInfo := internal.UserInfo{
		UserName:   req.GetName(),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := data.GetDb().Create(&userInfo)
	if ret.Error != nil {
		logger.Error(`AddUser db.Create error: %v`, ret.Error)
		return nil, ret.Error
	}

	return &library.AddUserResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
