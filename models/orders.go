package models

import (
	"github.com/jinzhu/gorm"
	u "../utils"
)

//a struct to rep PaymentInfo
type Orders struct {
	gorm.Model
	ID int `json:"id"`
	FirstLuggage int `json:"first_luggage"`
	C3TravelAPI int `json:"c3_travel_api"`
	MeMD int `json:"me_md"`
	ClientID int `json:"client_id"`
	UserID int `json:"user_id"`
}

func (orders *Orders) Create() (map[string] interface{}) {

	GetDB().Create(orders)

	if orders.ID <= 0 {
		return u.Message(false, "Failed to create orders, connection error.")
	}
	
	response := u.Message(true, "New orders has been added")
	response["orders"] = orders
	return response
}

func GetOrders(u uint) *Orders {

	p := &Orders{}
	GetDB().Table("payment_info").Where("id = ?", u).First(p)
	if p.ID <= 0 { //Orders not found!
		return nil
	}
	
	return p
}