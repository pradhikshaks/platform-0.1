package server

import (
	"fmt"
	"microservices/handler"
	"net/http"


	"github.com/gorilla/mux"
)

func RestServer() error {
	// _, b, _, _ := runtime.Caller(0)
	// basePath := filepath.Dir(filepath.Dir(b))
	database := handler.InfoFromEnv()
	handlers := handler.New(database)
	
	router := mux.NewRouter()

	router.HandleFunc("/data/{endpoint}",handlers.GetData).Methods("GET")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), router); err != nil {
		return err
	} else {
		fmt.Println("initialised")
	}
	fmt.Println("Debug")
	return nil
}