package constant

import "github.com/lib/pq"

var (
	PostgresUniqueViolationErr = pq.ErrorCode("23505")
)
