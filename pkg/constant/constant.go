package constant

import "github.com/lib/pq"

var (
	PostgresUniqueViolationErr = pq.ErrorCode("23505")
)

var (
	DefaultLimit = uint64(100)
	DefaultOffet = uint64(0)
)

var (
	DefaultBalance = 1000
)
