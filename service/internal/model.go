package internal

type BookInfo struct {
	Id         int64  `mysql:"id"`
	BookCode   string `mysql:"book_code"`
	BookName   string `mysql:"book_name"`
	BookCounts int32  `mysql:"book_counts"`
	UpdateTime int64  `mysql:"update_time"`
	UpdateBy   string `mysql:"update_by"`
}

type UserInfo struct {
	Id         int64  `mysql:"id"`
	UserName   string `mysql:"user_name"`
	UpdateTime int64  `mysql:"update_time"`
	UpdateBy   string `mysql:"update_by"`
}

type UserBorrowInfo struct {
	Id         int64  `mysql:"id"`
	UserId     int64  `mysql:"user_id"`
	BookCode   string `mysql:"book_code"`
	Status     int32  `mysql:"status"`
	UpdateTime int64  `mysql:"update_time"`
	UpdateBy   string `mysql:"update_by"`
}
