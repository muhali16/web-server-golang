package main

import (
	"encoding/json"
	. "fmt"
)

type User struct {
	Nama string `json:"nama"`
	Usia int    `json:"usia"`
}

func decodeJsonIntoStruct(data string) User {
	// var jsonString = data
	var jsonData = []byte(data)

	var user1 User

	var err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		Println(err.Error())
		return user1
	}

	return user1
}

func decodeJsonIntoMap(data string) map[string]interface{} {
	var jsonString = []byte(data)
	var jsonData = new(map[string]interface{})
	var err = json.Unmarshal(jsonString, jsonData)
	if err != nil {
		Println(err.Error())
		return *jsonData
	}
	return *jsonData
}

func decodeArrayJsonIntoArrayObject(data string) []User {
	var jsonData = []byte(data)
	var user1 []User
	var err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		Println(err.Error())
		return user1
	}
	return user1
}

func encodeObjectIntoJson(data *[]User) string {
	var jsonData, err = json.Marshal(data)
	if err != nil {
		Println(err.Error())
		return string(jsonData)
	}

	return string(jsonData)
}
