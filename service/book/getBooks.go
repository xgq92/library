package book

import (
	"context"
	"library/data"
	"library/grpc/library"
	"library/service/internal"

	logger "github.com/alecthomas/log4go"
)

func GetBooks(ctx context.Context, req *library.GetBooksReq) (*library.GetBooksResp, error) {
	// where := `"book_code" = ? AND "book_name" = ? "book_counts" <> 0`
	// if req.GetCode() == "" || req.GetName() == "" {
	// 	where = `"book_counts" <> 0`
	// }
	where := internal.BookInfo{BookCode: req.GetCode(), BookName: req.GetName()}
	rets := []internal.BookInfo{}
	var count int32
	// data.GetDb().Where(where, req.GetCode(), req.GetName()).Find(&rets).Count(&count)
	data.GetDb().Where(where).Find(&rets).Count(&count)
	if count != 0 {
		ret := data.GetDb().Limit(req.GetLimit()).Offset(req.GetOffset()).
			Where(where).Find(&rets)
		if ret.Error != nil {
			logger.Error(`DeleteBook db.Delete error: %v`, ret.Error)
			return nil, ret.Error
		}
	}

	return &library.GetBooksResp{
		Result:     &library.Result{Code: 0, Hint: "成功"},
		TotalCount: count,
		Datas:      BookInfoToPb(rets),
	}, nil
}

func BookInfoToPb(db []internal.BookInfo) []*library.BookList {
	rets := make([]*library.BookList, 0, len(db))
	for _, data := range db {
		ret := &library.BookList{
			BookCode:   data.BookCode,
			BookName:   data.BookName,
			BookCounts: data.BookCounts,
		}
		rets = append(rets, ret)
	}
	return rets
}
