package models

import (
	"github.com/jinzhu/gorm"
	u "../utils"
)

//a struct to rep PaymentInfo
type PaymentInfo struct {
	gorm.Model
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	ZIP string `json:"zip"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Gender string `json:"gender"`
}

func (paymentInfo *PaymentInfo) Create() (map[string] interface{}) {

	GetDB().Create(paymentInfo)

	if paymentInfo.ID <= 0 {
		return u.Message(false, "Failed to create paymentInfo, connection error.")
	}
	
	response := u.Message(true, "New paymentInfo has been added")
	response["paymentInfo"] = paymentInfo
	return response
}

func GetPaymentInfo(u uint) *PaymentInfo {

	p := &PaymentInfo{}
	GetDB().Table("payment_info").Where("id = ?", u).First(p)
	if p.ID <= 0 { //PaymentInfo not found!
		return nil
	}
	
	return p
}