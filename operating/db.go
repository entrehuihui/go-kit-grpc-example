package operating

import "github.com/jmoiron/sqlx"

// Orderoperating ..
type Orderoperating struct {
	db *sqlx.DB
}

// RetResponse 返回数据结构
type RetResponse struct {
	AllNum int
	Data   interface{}
}

// NewOrderoperating .
func NewOrderoperating(db *sqlx.DB) *Orderoperating {
	return &Orderoperating{
		db: db,
	}
}
