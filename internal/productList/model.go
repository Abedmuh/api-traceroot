package productlist

import (
	"time"

	"github.com/Abedmuh/api-traceroot/internal/products"
)

// main
//
//	type ProductList struct {
//		Id          string
//		Id_products string
//		Owner       string
//		TimeLimit   time.Time
//		Username    string
//		Password    string
//		Created_at  time.Time
//	}
type ProductList struct {
	Owner      string    `json:"owner" validate:"required"`
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Timelimit  time.Time `json:"time_limit" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Os         string    `json:"os" validate:"required"`
	Cpu        string    `json:"cpu" validate:"required"`
	Storage    string    `json:"storage" validate:"required"`
	Firewall   bool      `json:"firewall" validate:"required"`
	Selinux    string    `json:"selinux" validate:"required"`
	Location   string    `json:"location" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (ProductList) TableName() string {
	return "productlist"
}

// request
type ReqProdList struct {
	Id        string
	Products  products.Products
	Owner     string
	TimeLimit time.Time
}

// response
type ResProdList struct {
	Id          string
	Id_products string
	Owner       string
	TimeLimit   time.Time
	Username    string
	Password    string
	Created_at  time.Time
}
