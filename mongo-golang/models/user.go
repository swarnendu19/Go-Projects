package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `json:"name" bson: "name"`
	Gender string        `json: "gender" bson:"gender" `
	Age    int           `json: "age" bson:"age"`
}
