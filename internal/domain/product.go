package domain

// предпологается что продутов меньше чем пользователей (<100k)
// поэтому в качестве id используется uint64
type ProductID uint64

// возможно стоило разбить на Item + Price
type Product struct {
	ID    ProductID
	Name  string
	Price uint64
}
