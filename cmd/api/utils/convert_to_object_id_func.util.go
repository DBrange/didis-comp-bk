package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConvertToObjectIDFunc func(string) (*primitive.ObjectID, error)
