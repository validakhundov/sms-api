package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/validakhundov/sms-api/handler"
	"github.com/validakhundov/sms-api/properties"
	"net/http"
)

func main() {
	godotenv.Load("profiles/default.env")
	properties.LoadProperties()
	router := mux.NewRouter()
	handler.NewSmsHandler(router)
	http.ListenAndServe(":80", router)
}
