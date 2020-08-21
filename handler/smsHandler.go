package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/validakhundov/sms-api/client"
	"github.com/validakhundov/sms-api/model"
	"github.com/validakhundov/sms-api/service"
	"net/http"
)

type smsHandler struct {
	service service.SmsService
}

var smsService = service.SmsServiceImpl{
	AzercellCabinetClient: &client.AzercellCabinetClientImpl{},
}

func NewSmsHandler(router *mux.Router) *mux.Router {
	h := &smsHandler{service: &smsService}
	router.HandleFunc("/v1/send", h.sendSms).Methods("POST")
	return router
}

func (h *smsHandler) sendSms(w http.ResponseWriter, r *http.Request) {
	var sms model.Sms
	json.NewDecoder(r.Body).Decode(&sms)
	resp, err := smsService.SendSms(sms)
	if err != nil {
		http.Error(w, err.Error(), 200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
