package user

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"

	logger "github.com/alecthomas/log4go"
)

func DeleteUser(ctx context.Context, req *library.DeleteUserReq) (*library.DeleteUserResp, error) {
	ret := data.GetDb().Where("id = ?", req.GetUserId()).Delete(&internal.UserInfo{})
	if ret.Error != nil {
		logger.Error(`DeleteUser db.Delete error: %v`, ret.Error)
		return nil, ret.Error
	}
	return &library.DeleteUserResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
