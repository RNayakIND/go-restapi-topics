package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type EmployeeDetails struct {
	Name           string         `json:name`
	AddressDetails AddressDetails `json:address`
	Location       string         `json:location`
}

type AddressDetails struct {
	Street string      `json:street`
	Sector interface{} `json:sector`
}

var employeeDetails []EmployeeDetails

func getAllEmployee(response http.ResponseWriter, request *http.Request) {
	log.Print("Login ....!!!!")
	json.NewEncoder(response).Encode(employeeDetails)

}
func getEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	param := (mux.Vars(request))

	for index, employeeRange := range employeeDetails {
		if employeeRange.AddressDetails.Sector == param["sector"] {
			employeeDetails = append(employeeDetails[:index], employeeDetails[index+1:]...)
			fmt.Print("Employee Found")
			json.NewEncoder(response).Encode(employeeDetails)
		}

	}
	json.NewEncoder(response).Encode(employeeDetails)

}

func addEmployee(response http.ResponseWriter, request *http.Request) {
	var addEmployeeNew EmployeeDetails
	response.Header().Set("Content-Type", "application/json")
	json.NewDecoder(request.Body).Decode(&addEmployeeNew)
	employeeDetails = append(employeeDetails, addEmployeeNew)
	for _, employee := range employeeDetails {
		fmt.Println("Employee Details::::", employee.AddressDetails.Sector)
	}
	json.NewEncoder(response).Encode(employeeDetails)
}

func main() {

	router := mux.NewRouter()
	employeeDetails = append(employeeDetails,
		EmployeeDetails{Name: "Ranjan ", Location: "Hyderabad", AddressDetails: AddressDetails{Street: "PATIA", Sector: "33"}},
		EmployeeDetails{Name: "Rakesh ", Location: "Bhubabneswar", AddressDetails: AddressDetails{Street: "ShikaharChandi", Sector: "33"}},
		EmployeeDetails{Name: "Mohit ", Location: "Cuttack", AddressDetails: AddressDetails{Street: "CuttackChandi", Sector: "30"}},
	)
	log.Print("EmployeeDetails::::", employeeDetails)
	router.HandleFunc("/getAllEmployee", getAllEmployee).Methods("GET")
	router.HandleFunc("/getEmployee/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/addEmployee", addEmployee).Methods("POST")
	log.Fatal(http.ListenAndServe(":9111", router))

}
