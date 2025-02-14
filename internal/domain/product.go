package domain

// предпологается что продутов меньше чем пользователей
type ProductID uint64

type Product struct {
	ID    ProductID
	Name  string
	Price uint64
}
