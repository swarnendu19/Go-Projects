package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/swarnendu19/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}
