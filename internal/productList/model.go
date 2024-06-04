package productlist

import "time"

//main
type ProductList struct {
	Id          string
	Id_products string
	Owner       string
	TimeLimit   time.Time
	Username    string
	Password    string
	Created_at  time.Time
}

//request
type ReqProdList struct {
	Id          string
	Id_products string
	Owner       string
	TimeLimit   time.Time
}

//response
type ResProdList struct {
	Id          string
	Id_products string
	Owner       string
	TimeLimit   time.Time
	Username    string
	Password    string
	Created_at  time.Time
}
