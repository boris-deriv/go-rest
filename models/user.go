package models

type Address struct {
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
	Street  string `json:"street" bson:"street"`
}

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     int     `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
}
