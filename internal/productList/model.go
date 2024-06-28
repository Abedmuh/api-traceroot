package productlist

import (
	"time"
)

type ProductList struct {
	Owner     string    `json:"owner" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Timelimit time.Time `json:"time_limit" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Os        string    `json:"os" validate:"required"`
	Cpu       int32    `json:"cpu" validate:"required"`
	Ram       int64   `json:"ram" validate:"required"`
	Storage   int64    `json:"storage" validate:"required"`
	Firewall  bool      `json:"firewall" validate:"required"`
	Selinux   string    `json:"selinux" validate:"required"`
	Location  string    `json:"location" validate:"required"`
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

type OsDetails struct {
	Name        string
	Version     string
	GuestId 	string
	Location 	string
}

var OsMap = map[string]OsDetails{
	"Ubuntu": {
        Name:        "Ubuntu",
        Version:     "18.04",
        GuestId:     "ubuntu-18.04",
        Location:     "US West",
    },
    "CentOS": {
        Name:        "CentOS",
        Version:     "7.6",
        GuestId:     "centos7_64Guest",
        Location:     "US East",
    },
    // Add more Os details here
}