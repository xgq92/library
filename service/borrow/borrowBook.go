package borrow

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func BorrowBook(ctx context.Context, req *library.BorrowBookReq) (*library.BorrowBookResp, error) {
	// 开始事务
	tx := data.GetDb().Begin()
	userInfo := internal.UserBorrowInfo{
		UserId:     req.GetUserId(),
		BookCode:   req.GetBookCode(),
		Status:     int32(library.Status_Borrowing.Number()),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := tx.Create(&userInfo)
	if ret.Error != nil {
		logger.Error(`BorrowBook db.Create error: %v`, ret.Error)
		tx.Rollback()
		return nil, ret.Error
	}

	bookData := internal.BookInfo{}
	err := tx.Where("book_code = ?", req.GetBookCode()).First(&bookData).Error
	if err != nil {
		logger.Error(`BorrowBook bookData db.First error: %v`, err)
		tx.Rollback()
		return nil, err
	}
	if bookData.BookCounts > 0 {
		bookData.BookCounts--
		err := tx.Model(&bookData).Where("book_code = ?", req.GetBookCode()).
			Update("book_counts", bookData.BookCounts).Error
		if err != nil {
			logger.Error(`BorrowBook db.Update book_code error: %v`, err)
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	return &library.BorrowBookResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
