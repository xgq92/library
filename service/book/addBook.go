package book

import (
	"context"
	"fmt"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func AddBook(ctx context.Context, req *library.AddBookReq) (*library.AddBookResp, error) {
	bookCode, err := BulidBookCode(ctx)
	if err != nil {
		logger.Error(`AddBook BulidBookCode error: %v`, err)
		return nil, err
	}
	bookInfo := &internal.BookInfo{
		BookCode:   bookCode,
		BookName:   req.GetName(),
		BookCounts: req.GetCounts(),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := data.GetDb().Create(bookInfo)
	if ret.Error != nil {
		logger.Error(`AddBook db.Create error: %v`, ret.Error)
		return nil, ret.Error
	}

	return &library.AddBookResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}

func BulidBookCode(ctx context.Context) (string, error) {
	dateStr := time.Now().Format("200601")
	id, err := getBookCount(ctx, dateStr)
	if err != nil {
		return "", err
	}
	idStr := fmt.Sprintf("%0.3d", id)
	return fmt.Sprintf("%s%s%s", "BK", dateStr, idStr), nil
}

func getBookCount(ctx context.Context, dateStr string) (int64, error) {
	ret := data.GetRedis().Incr(ctx, "books_count"+dateStr)
	ret.Val()
	if ret.Err() != nil {
		logger.Error(`GetBookCount GetRedis().Incr error: %v`, ret.Err())
		return 0, ret.Err()
	}
	if ret.Val() == 1 {
		data.GetRedis().Expire(ctx, "books_count"+dateStr, time.Hour*24*32)
	}
	return ret.Val(), nil
}
