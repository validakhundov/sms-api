package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/validakhundov/sms-api/client"
	"github.com/validakhundov/sms-api/model"
	"testing"
)

var (
	mockClient = client.AzercellCabinetClientMock{}
	service    = SmsServiceImpl{AzercellCabinetClient: &mockClient}
)

func TestSmsServiceSendSmsSuccess(t *testing.T) {
	response := new(model.Response)
	response.Code = 0
	response.Message = "Message sent successfully!"
	response.Type = "free"
	sms := new(model.Sms)
	sms.Text = "test"
	sms.Number = "0101010"

	result, err := service.SendSms(*sms)

	assert.Nil(t, err)
	assert.Equal(t, response.Code, result.Code)
	assert.Equal(t, response.Message, result.Message)
	assert.Equal(t, response.Type, result.Type)
}
