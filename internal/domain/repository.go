package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       string `bson:"id"`
	Fio      *Fio   `bson:"fio"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Driver   bool   `bson:"driver"`
	Customer bool   `bson:"customer"`
}

type Fio struct {
	Name       string `bson:"name"`
	Surname    string `bson:"surname"`
	Patronymic string `bson:"patronymic"`
}

type FormInput struct {
	UserID      string    `bson:"user_id"`
	WhereFrom   string    `bson:"where_from"`
	WhereTo     string    `bson:"where_to"`
	Description string    `bson:"description"`
	Time        time.Time `bson:"time"`
}

type TestDataRepo struct {
	ID         primitive.ObjectID `bson:"id"`
	DriverName string             `bson:"driver_name"`
	CarsName   string             `bson:"cars_name"`
}
