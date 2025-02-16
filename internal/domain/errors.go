package domain

import "errors"

// Users
var (
	ErrUserNotFound = errors.New("user not found")
)

// Products
var (
	ErrProductNotFound = errors.New("product not found")
)

// Trasfers
var (
	ErrZeroAmount          = errors.New("amount must be greater than zero")
	ErrSameUser            = errors.New("cannot send coins to yourself")
	ErrMissingFromUser     = errors.New("sender user ID must be specified")
	ErrMissingToUser       = errors.New("recipient user ID must be specified")
	ErrInsufficientBalance = errors.New("insufficient balance")
)
