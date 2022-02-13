package store

import (
	"github.com/ishank838/go-users-grpc/helper"
	users "github.com/ishank838/go-users-grpc/proto/user"
)

type sampleUser struct {
	id        int64
	firstname string
	city      string
	phone     string
	height    float32
	married   bool
}

var UserDb = []users.User{}

func InituserData() {
	data := []sampleUser{
		{id: 1, firstname: "Abey", city: "Houxiang", phone: "4824120772", height: 5.7, married: false},
		{id: 2, firstname: "Mathilda", city: "Avellaneda", phone: "1729774291", height: 5.9, married: false},
		{id: 3, firstname: "Cyndia", city: "Dahuang", phone: "7827784404", height: 5.8, married: false},
		{id: 4, firstname: "Rudolf", city: "Kasamatsuchō", phone: "6881851018", height: 5.4, married: false},
		{id: 5, firstname: "Noella", city: "Rukaj", phone: "5098942815", height: 5.2, married: true},
		{id: 6, firstname: "Cass", city: "Prévost", phone: "8405606760", height: 5.8, married: true},
		{id: 7, firstname: "Myrtice", city: "Pingfeng", phone: "5886662614", height: 6.0, married: false},
		{id: 8, firstname: "Pavia", city: "Demak", phone: "7235725175", height: 5.10, married: true},
		{id: 9, firstname: "Bevin", city: "Santa Ana", phone: "3616661321", height: 5.2, married: false},
		{id: 10, firstname: "Chase", city: "Conchal", phone: "8021816836", height: 5.9, married: false},
	}

	for i, _ := range data {
		UserDb = append(UserDb, users.User{Id: &data[i].id,
			FirstName: &data[i].firstname, City: &data[i].city, Phone: &data[i].phone, Height: helper.GetFormattedHeight(data[i].height),
			Married: &data[i].married,
		})
	}
}

func GetUserById(id int64) *users.User {
	for _, u := range UserDb {
		if *u.Id == id {
			return &u
		}
	}
	return nil
}

func GetAllUsers() []users.User {
	return UserDb
}
