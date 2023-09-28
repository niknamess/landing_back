package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestData struct {
	Data string
}

type TestDataService struct {
	ID         primitive.ObjectID
	DriverName string
	CarsName   string
}
