package domain

// uint64 достаточно для хранения
// 100k уникальных пользователей
type UserID uint64

type User struct {
	ID       UserID
	Username string
}
