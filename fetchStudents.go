package main

import (
	"bytes"
	"encoding/json"
	. "fmt"
	"net/http"
	"net/url"
)

const baseURL = "http://localhost:8000"

type Student struct {
	Name  string `json:"nama"`
	Id    int    `json:"nis"`
	Grade int    `json:"nilai"`
}

func fetchStudents() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var students []Student

	var request, err1 = http.NewRequest("GET", baseURL+"/students", nil)
	if err1 != nil {
		return nil, err1
	}

	var response, err2 = client.Do(request)
	if err2 != nil {
		return nil, err2
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func fetchStudent(id string) (Student, error) {
	var err error
	var student Student
	var client = &http.Client{}

	var param = url.Values{}
	param.Set("id", id)
	var payload = bytes.NewBufferString(param.Encode())

	var request, err1 = http.NewRequest("POST", baseURL+"/student", payload)
	if err1 != nil {
		Println("Error while requesting data.")
		return student, err1
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response, err2 = client.Do(request)
	if err2 != nil {
		Println("Error while giving response the data.")
		return student, err2
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&student)
	if err != nil {
		return student, err
	}

	return student, nil
}
