package models

import (
	u "github.com/Lockwarr/POSAR/utils"
	"github.com/jinzhu/gorm"
)

//a struct to rep PaymentInfo
type Orders struct {
	gorm.Model
	FirstLuggage int `json:"first_luggage"`
	C3TravelAPI  int `json:"c3_travel_api"`
	MeMD         int `json:"me_md"`
	ClientID     int `json:"client_id"`
	UserID       int `json:"user_id"`
}

func (orders *Orders) Create() map[string]interface{} {

	GetDB().Create(orders)

	if orders.ID <= 0 {
		return u.Message(false, "Failed to create orders, connection error.")
	}

	response := u.Message(true, "New orders has been added")
	response["orders"] = orders
	return response
}

func GetOrdersForClient(clientID int) *Orders {

	p := &Orders{}
	GetDB().Table("payment_info").Where("client_id = ?", clientID).First(p)
	if p.ID <= 0 { //Orders not found!
		return nil
	}

	return p
}

func GetOrdersForUser(userID int) *Orders {

	p := &Orders{}
	GetDB().Table("payment_info").Where("user_id = ?", userID).First(p)
	if p.ID <= 0 { //Orders not found!
		return nil
	}

	return p
}
