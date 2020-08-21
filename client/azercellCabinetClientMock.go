package client

import (
	"github.com/stretchr/testify/mock"
	"github.com/validakhundov/sms-api/model"
)

type AzercellCabinetClientMock struct {
	mock.Mock
}

func (c *AzercellCabinetClientMock) GetToken() (string, error) {
	return "f304a2da-2664-41e8-a", nil
}

func (c *AzercellCabinetClientMock) SendSms(t string, s model.AzercellSmsBody) (*model.AzercellSmsResponse, error) {
	response := new(model.AzercellSmsResponse)
	stringResponse := new(model.AzercellSmsStringResponse)
	stringResponse.ResponseMessage = "ok.free"
	response.StringResponse = *stringResponse
	return response, nil
}
