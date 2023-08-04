package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MedicalPatients struct {
	ID                  int                  `json:"id,omitempty`
	Firstname           string               `json:"firstname,omitempty`
	Lastname            string               `json:"lastname,omitempty`
	Email               string               `json:"email,omitempty`
	MedicalAppointments *MedicalAppointments `json:"medicalappointments,omitempty`
}
type MedicalAppointments struct {
	ID              int    `json:"id,omitempty`
	Description     string `json:"description,omitempty`
	Appointmentdate string `json:"appointmentdate,omitempty`
	Appointmenttime string `json:"appointmenttime,omitempty`
}

func GetPople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

var people []MedicalPatients

func main() {
	router := mux.NewRouter()
	fmt.Println("Backend is running...")

	people = append(people,
		MedicalPatients{
			ID:                  1,
			Firstname:           "John",
			Lastname:            "Doe",
			Email:               "johndoe@nothing.com",
			MedicalAppointments: &MedicalAppointments{ID: 2, Description: "Medical Appointments", Appointmentdate: "04/08/2023", Appointmenttime: "17:00"}})

	router.HandleFunc("/MedicalPatients", GetPople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
