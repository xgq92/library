package book

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func UpdateBook(ctx context.Context, req *library.UpdateBookReq) (*library.UpdateBookResp, error) {
	booKInfo := internal.BookInfo{
		BookCode:   req.GetBookCode(),
		BookName:   req.GetBookName(),
		BookCounts: req.GetBookCounts(),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := data.GetDb().Model(&booKInfo).Updates(booKInfo)
	if ret.Error != nil {
		logger.Error(`DeleteBook db.Updates error: %v`, ret.Error)
		return nil, ret.Error
	}
	return &library.UpdateBookResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
