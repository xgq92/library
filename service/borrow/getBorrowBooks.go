package borrow

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"

	logger "github.com/alecthomas/log4go"
)

type RetUserBorrowInfo struct {
	Id       int64  `mysql:"id"`
	UserId   int64  `mysql:"user_id"`
	BookCode string `mysql:"book_code"`
	Status   int32  `mysql:"status"`
	UserName string `mysql:"user_name"`
	BookName string `mysql:"book_name"`
}

func GetBorrowBooks(ctx context.Context, req *library.GetBorrowBooksReq) (*library.GetBorrowBooksResp, error) {
	where := internal.UserBorrowInfo{UserId: req.GetUserId()}
	dbModel := []internal.UserBorrowInfo{}
	rets := []RetUserBorrowInfo{}
	var count int32
	data.GetDb().Where(where).Find(&dbModel).Count(&count)
	if count != 0 {
		query := getSql(req.GetUserId())
		values := []interface{}{req.GetLimit(), req.GetOffset()}
		if req.GetUserId() > 0 {
			values = []interface{}{req.GetUserId(), req.GetLimit(), req.GetOffset()}
		}
		data.GetDb().Raw(query, values...).Scan(&rets)
		ret := data.GetDb().Raw(query, values...).Scan(&rets)
		if ret.Error != nil {
			logger.Error(`GetBorrowBooks db.Raw error: %v`, ret.Error)
			return nil, ret.Error
		}
	}

	return &library.GetBorrowBooksResp{
		Result:     &library.Result{Code: 0, Hint: "成功"},
		TotalCount: count,
		Datas:      UserBorrowInfoToPb(rets),
	}, nil
}

func UserBorrowInfoToPb(db []RetUserBorrowInfo) []*library.BorrowList {
	rets := make([]*library.BorrowList, 0, len(db))
	for _, data := range db {
		ret := &library.BorrowList{
			UserId:   data.UserId,
			BookCode: data.BookCode,
			UserName: data.UserName,
			BookName: data.BookName,
			Status:   library.Status(data.Status),
		}
		rets = append(rets, ret)
	}
	return rets
}

func getSql(userid int64) string {
	where := ""
	if userid > 0 {
		where = "WHERE a.user_id = ? "
	}
	query := `SELECT a.id,a.user_id,a.book_code,a.status,b.user_name,c.book_name 
		FROM user_borrow_infos a 
		INNER JOIN user_infos b ON a.user_id = b.id
		INNER JOIN book_infos c ON a.book_code = c.book_code ` +
		where +
		`LIMIT ? OFFSET ?`
	return query
}
