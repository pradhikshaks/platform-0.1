package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//	type EmployeeData struct {
//		name string
//		id string
//	}
type EmployeeDetials struct {
	Name string
	Id   string
	Role string
}

type ProjectDetials struct {
	Name    string
	Id      string
	Project string
}

type EmployeeData struct{
	Data []EmployeeDetials
}

type ProjectData struct{
	Data []ProjectDetials
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (h handler) GetData(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	vars := mux.Vars(r)
	endpoint := vars["endpoint"]
	defer r.Body.Close()
	fmt.Println("end", endpoint)
	if endpoint == "employee" {
		var resData []EmployeeDetials
		h.DB.Raw("SELECT * FROM emp").Scan(&resData)
		fmt.Println("res", resData)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := EmployeeData{
			Data: resData,
		}
		json.NewEncoder(w).Encode(data)
	} else if endpoint == "project" {
		var resData []ProjectDetials
		h.DB.Raw("SELECT * FROM projects").Scan(&resData)
		fmt.Println("res", resData)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := ProjectData{
			Data: resData,
		}
		json.NewEncoder(w).Encode(data)
	} else {

		resData := []EmployeeDetials{}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := EmployeeData{
			Data: resData,
		}
		json.NewEncoder(w).Encode(data)
		
	}

	
}

// docker exec -it postgresql-database-1 bash

// psql -U vibav -d db
