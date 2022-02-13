package store

import (
	users "github.com/ishank838/go-users-grpc/proto/user"
)

var userDb = []users.User{
	{Id: 1, FirstName: "Abey", City: "Houxiang", Phone: "482-412-0772", Height: 5.7, Married: false},
	{Id: 2, FirstName: "Mathilda", City: "Avellaneda", Phone: "172-977-4291", Height: 5.9, Married: false},
	{Id: 3, FirstName: "Cyndia", City: "Dahuang", Phone: "782-778-4404", Height: 5.8, Married: false},
	{Id: 4, FirstName: "Rudolf", City: "Kasamatsuchō", Phone: "688-185-1018", Height: 5.4, Married: false},
	{Id: 5, FirstName: "Noella", City: "Rukaj", Phone: "509-894-2815", Height: 5.2, Married: true},
	{Id: 6, FirstName: "Cass", City: "Prévost", Phone: "840-560-6760", Height: 5.8, Married: true},
	{Id: 7, FirstName: "Myrtice", City: "Pingfeng", Phone: "588-666-2614", Height: 6.0, Married: false},
	{Id: 8, FirstName: "Pavia", City: "Demak", Phone: "723-572-5175", Height: 5.10, Married: true},
	{Id: 9, FirstName: "Bevin", City: "Santa Ana", Phone: "361-666-1321", Height: 5.2, Married: false},
	{Id: 10, FirstName: "Chase", City: "Conchal", Phone: "802-181-6836", Height: 5.9, Married: false},
}

func GetUserById(id int64) *users.User {
	for _, u := range userDb {
		if u.Id == id {
			return &u
		}
	}
	return nil
}

func GetAllUsers() *[]users.User {
	return &userDb
}
