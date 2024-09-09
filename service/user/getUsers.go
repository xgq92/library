package user

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"

	logger "github.com/alecthomas/log4go"
)

func GetUsers(ctx context.Context, req *library.GetUsersReq) (*library.GetUsersResp, error) {
	// where := `"id" = ? AND "user_name" = ?`
	where := internal.UserInfo{Id: req.GetUserId(), UserName: req.GetName()}
	rets := []internal.UserInfo{}
	var count int32
	// err := data.GetDb().Where(where, req.GetUserId(), req.GetName()).Find(&rets).Count(&count).Error
	err := data.GetDb().Where(where).Find(&rets).Count(&count).Error
	if err != nil {
		logger.Error(`GetUsers db.Find.count error: %v`, err)
		return nil, err
	}
	if count != 0 {
		ret := data.GetDb().Limit(req.GetLimit()).Offset(req.GetOffset()).
			Where(where).Find(&rets)
		if ret.Error != nil {
			logger.Error(`GetUsers db.Find error: %v`, ret.Error)
			return nil, ret.Error
		}
	}

	return &library.GetUsersResp{
		Result:     &library.Result{Code: 0, Hint: "成功"},
		TotalCount: count,
		Datas:      UserInfoToPb(rets),
	}, nil
}

func UserInfoToPb(db []internal.UserInfo) []*library.UserList {
	rets := make([]*library.UserList, 0, len(db))
	for _, data := range db {
		ret := &library.UserList{
			UserId:   data.Id,
			UserName: data.UserName,
		}
		rets = append(rets, ret)
	}
	return rets
}
