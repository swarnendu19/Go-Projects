package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/swarnendu19/mongo-golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Retrieve the "id" parameter from the request URL using the ByName method from the p object (typically, p is an instance of httprouter.Params).
	id := p.ByName("id")

	//Check if the retrieved id is a valid hexadecimal representation of a BSON ObjectId. If not, write a 404 Not Found status to the response header.
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	//Convert the valid hexadecimal id to a BSON ObjectId.

	oid := bson.ObjectIdHex(id)

	//Create a new instance of the User struct from the models package.
	users := models.User{}
	//Try to find a user in the "users" collection of the "mongo-golang" database with the specified ObjectId (oid). If an error occurs (e.g., the user is not found), write a 404 Not Found status to the response header and return from the handler.
	err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&users)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//Marshal the User struct u into a JSON byte slice (uj). If there is an error during marshalling, print the error.

	uj, err := json.Marshal(users)
	if err != nil {
		return
	}
	//Set the "Content-Type" header of the response to "application/json".
	w.Header().Set("Content-Type", "application/json")

	//Write a 200 OK status to the response header.

	w.WriteHeader(http.StatusOK)

	//Write the marshalled JSON (uj) to the response body.

	fmt.Println("User", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
