package service

import (
	"github.com/validakhundov/sms-api/client"
	"github.com/validakhundov/sms-api/model"
	"strings"
)

type SmsService interface {
	SendSms(sms model.Sms) (*model.Response, error)
}

type SmsServiceImpl struct {
	AzercellCabinetClient client.AzercellCabinetClient
}

func (s *SmsServiceImpl) SendSms(sms model.Sms) (*model.Response, error) {
	response := new(model.Response)
	token, _ := s.AzercellCabinetClient.GetToken()
	body := new(model.AzercellSmsBody)
	body.Date = "now"
	body.Message = sms.Text
	body.SendTo = sms.Number
	resp, err := s.AzercellCabinetClient.SendSms(token, *body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(resp.StringResponse.ResponseMessage, "ok") {
		response.Code = 0
		response.Message = "Message sent successfully!"
		response.Type = strings.Replace(resp.StringResponse.ResponseMessage, "ok.", "", 1)
	} else {
		response.Code = -1
		response.Message = strings.Replace(resp.StringResponse.ResponseMessage, "err.", "", 1)
	}
	return response, err
}