package models

import (
	u "github.com/Lockwarr/POSAR/utils"
	"github.com/jinzhu/gorm"
)

//FirstLuggage - a struct to rep FirstLuggage model
type FirstLuggage struct {
	gorm.Model
	ID            int    `json:"id"`
	CountryFrom   string `json:"country_from"`
	CountryTo     string `json:"country_to"`
	Return        bool   `json:"return"`
	PaymentInfoID int    `json:"payment_info_id"`
}

//Create - creates new luggage quote
func (firstLuggage *FirstLuggage) Create() map[string]interface{} {

	GetDB().Create(firstLuggage)

	if firstLuggage.ID <= 0 {
		return u.Message(false, "Failed to create firstLuggage, connection error.")
	}

	response := u.Message(true, "New firstLuggage has been added")
	response["firstLuggage"] = firstLuggage
	return response
}

//GetFirstLuggage - retrieves luggage quote by index
func GetFirstLuggage(u uint) *FirstLuggage {

	f := &FirstLuggage{}
	GetDB().Table("first_luggage").Where("id = ?", u).First(f)
	if f.ID <= 0 { //FirstLuggage not found!
		return nil
	}

	return f
}
