package book

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"

	logger "github.com/alecthomas/log4go"
)

func DeleteBook(ctx context.Context, req *library.DeleteBookReq) (*library.DeleteBookResp, error) {
	ret := data.GetDb().Where("book_code = ?", req.GetBookCode()).Delete(&internal.BookInfo{})
	if ret.Error != nil {
		logger.Error(`DeleteBook db.Delete error: %v`, ret.Error)
		return nil, ret.Error
	}
	return &library.DeleteBookResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
