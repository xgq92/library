package service

import (
	"context"
	"library/grpc/library"
	"library/service/book"
	"library/service/borrow"
	"library/service/user"
)

type Service struct{}

func (s *Service) SayHello(_ context.Context, req *library.HelloRequest) (*library.HelloResponse, error) {
	return &library.HelloResponse{Message: req.Name + " world"}, nil
}

// book
func (s *Service) AddBook(ctx context.Context, req *library.AddBookReq) (*library.AddBookResp, error) {
	return book.AddBook(ctx, req)
}

func (s *Service) GetBooks(ctx context.Context, req *library.GetBooksReq) (*library.GetBooksResp, error) {
	return book.GetBooks(ctx, req)
}

func (s *Service) UpdateBook(ctx context.Context, req *library.UpdateBookReq) (*library.UpdateBookResp, error) {
	return book.UpdateBook(ctx, req)
}

func (s *Service) DeleteBook(ctx context.Context, req *library.DeleteBookReq) (*library.DeleteBookResp, error) {
	return book.DeleteBook(ctx, req)
}

// user
func (s *Service) AddUser(ctx context.Context, req *library.AddUserReq) (*library.AddUserResp, error) {
	return user.AddUser(ctx, req)
}

func (s *Service) GetUsers(ctx context.Context, req *library.GetUsersReq) (*library.GetUsersResp, error) {
	return user.GetUsers(ctx, req)
}

func (s *Service) UpdateUser(ctx context.Context, req *library.UpdateUserReq) (*library.UpdateUserResp, error) {
	return user.UpdateUser(ctx, req)
}

func (s *Service) DeleteUser(ctx context.Context, req *library.DeleteUserReq) (*library.DeleteUserResp, error) {
	return user.DeleteUser(ctx, req)
}

// borrow
func (s *Service) BorrowBook(ctx context.Context, req *library.BorrowBookReq) (*library.BorrowBookResp, error) {
	return borrow.BorrowBook(ctx, req)
}

func (s *Service) ReturnBook(ctx context.Context, req *library.ReturnBookReq) (*library.ReturnBookResp, error) {
	return borrow.ReturnBook(ctx, req)
}

func (s *Service) GetBorrowBooks(ctx context.Context, req *library.GetBorrowBooksReq) (*library.GetBorrowBooksResp, error) {
	return borrow.GetBorrowBooks(ctx, req)
}
