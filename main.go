package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Lastname  string `json:lastname`
	Firstname string `json:firstname`
	Age       int    `json:age`
	Email     string `json:email`
	Phone     string `json:phone`
}

func main() {
	jsonFromApi := `
	[
		{
			"lastname": "strange",
			"firstname": "stephen",
			"age": 38,
			"email": "stephen.strange@exemple.com",
			"phone":"0123456789"
		},
		{
			"lastname": "parker",
			"firstname": "peter",
			"age": 25,
			"email": "peter.parker@exemple.com",
			"phone":"1234567890"
		}
	]`

	var users []User

	err := json.Unmarshal([]byte(jsonFromApi), &users)
	if err != nil {
		fmt.Println("Erreur unmarshalling json:", err)
	}

	fmt.Printf("json: %v\n", users)

	//-------------------------------

	var myStruct []User

	var user_one User
	user_one.Lastname = "wayne"
	user_one.Firstname = "bruce"
	user_one.Age = 34
	user_one.Email = "bruce.wayne@exemple.com"
	user_one.Phone = "2345678901"

	var user_two User
	user_two.Lastname = "wayne"
	user_two.Firstname = "bruce"
	user_two.Age = 34
	user_two.Email = "bruce.wayne@exemple.com"
	user_two.Phone = "2345678901"

	myStruct = append(myStruct, user_one)
	myStruct = append(myStruct, user_two)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}
	fmt.Println(string(jsonFromSlice))
}
