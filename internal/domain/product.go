package domain

// предпологается что продутов меньше чем пользователей, поэтому id uint64
type ProductID uint64

type Product struct {
	ID    ProductID
	Name  string
	Price uint64
}
