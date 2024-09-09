package borrow

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"
	"time"

	logger "github.com/alecthomas/log4go"
)

func ReturnBook(ctx context.Context, req *library.ReturnBookReq) (*library.ReturnBookResp, error) {
	// 开始事务
	tx := data.GetDb().Begin()
	userBorrowInfo := internal.UserBorrowInfo{
		UserId:     req.GetUserId(),
		BookCode:   req.GetBookCode(),
		Status:     int32(library.Status_BorrowReturn),
		UpdateTime: time.Now().Unix(),
		UpdateBy:   "管理员",
	}
	ret := tx.Model(&userBorrowInfo).Updates(userBorrowInfo)
	if ret.Error != nil {
		logger.Error(`ReturnBook db.Updates error: %v`, ret.Error)
		tx.Rollback()
		return nil, ret.Error
	}

	bookData := internal.BookInfo{}
	err := tx.Where("book_code = ?", req.GetBookCode()).First(&bookData).Error
	if err != nil {
		logger.Error(`ReturnBook bookData db.First error: %v`, ret.Error)
		tx.Rollback()
		return nil, err
	}
	if bookData.Id > 0 {
		bookData.BookCounts++
		err := tx.Model(&bookData).Where("book_code = ?", req.GetBookCode()).
			Update("book_counts", bookData.BookCounts).Error
		if err != nil {
			logger.Error(`ReturnBook db.Update book_code error: %v`, err)
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	return &library.ReturnBookResp{
		Result: &library.Result{Code: 0, Hint: "成功"},
	}, nil
}
