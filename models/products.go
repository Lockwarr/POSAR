package models

import (
	u "github.com/Lockwarr/POSAR/utils"
	"github.com/jinzhu/gorm"
)

//a struct to rep selected products
type Products struct {
	gorm.Model
	ClientID            int  `json:"client_id", gorm:"foreignkey:ClientID"`
	FirstLuggagePresent bool `json:"first_luggage_present"`
	C3TravelApiPresent  bool `json:"c3_travel_api_present"`
	MeMDPresent         bool `json:"me_md_present"`
}

func (products *Products) Update() map[string]interface{} {
	//if GetClient(uint(products.ClientID)) == nil {
	//	return u.Message(false, "Client with this ID doesn't exist.")
	//}

	GetDB().Save(&products)

	if products.ID <= 0 {
		return u.Message(false, "Failed to create products, connection error.")
	}

	response := u.Message(true, "New products has been added")
	response["products"] = products
	return response
}

func GetProducts(u uint) *Products {

	p := &Products{}
	GetDB().Table("products").Where("id = ?", u).First(p)
	if p.ClientID <= 0 { //Products for client not found!
		return nil
	}

	return p
}
