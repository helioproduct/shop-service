package domain

import "errors"

var (
	ErrInternalError = errors.New("internal service error")
)

// Users
var (
	ErrPasswordRequired = errors.New("password required")
	ErrUsernameRequired = errors.New("username required")
	ErrUserNotFound     = errors.New("user not found")
	ErrUserExists       = errors.New("user already exists")
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

// auth
var (
	ErrParsingToken         = errors.New("error parsing token")
	ErrInvalidToken         = errors.New("invalid token")
	ErrExpiredToken         = errors.New("token expired")
	ErrTokenSessionMismatch = errors.New("session data mismatch")
	ErrInvalidCredentials   = errors.New("invalid credentials")
)
