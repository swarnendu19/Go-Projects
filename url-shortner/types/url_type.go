package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// Make a type for Client What exactly client is going to send us

type shortUrlDB struct {
	longUrl string `json:"long_url"`
}

//Create a type for Database side What exactly the mongo db databse will store and how to store

type urlDB struct {
	ID        primitive.ObjectID `bson: "_id", omitempty `
	UrlCode   string             `bson: "url_code"`
	longUrl   string             `bson: "long_url"`
	shortUrl  string             `bson: "short_url"`
	createdAt int64              `bson:"created_at"`
	ExpiredAt int64              `bson:"expired_at"`
}
