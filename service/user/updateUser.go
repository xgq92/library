package user

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func UpdateUser(ctx context.Context, req *library.UpdateUserReq) (*library.UpdateUserResp, error) {
	userInfo := internal.UserInfo{
		Id:         req.GetUserId(),
		UserName:   req.GetUserName(),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := data.GetDb().Model(&userInfo).Updates(userInfo)
	if ret.Error != nil {
		logger.Error(`DeleteBook db.Updates error: %v`, ret.Error)
		return nil, ret.Error
	}
	return &library.UpdateUserResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
