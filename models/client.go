package models

import (
	"github.com/jinzhu/gorm"
	u "../utils"
)

//a struct to rep client
type Client struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
}

func (client *Client) Create() (map[string] interface{}) {

	GetDB().Create(client)

	if client.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}
	
	response := u.Message(true, "New client has been added")
	response["client"] = client
	return response
}

func GetUser(u uint) *Client {

	c := &Client{}
	GetDB().Table("client").Where("id = ?", u).First(c)
	if c.Name == "" { //Client not found!
		return nil
	}

	return c
}