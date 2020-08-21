package client

import (
	"bytes"
	"encoding/json"
	"github.com/validakhundov/sms-api/model"
	"github.com/validakhundov/sms-api/properties"
	"io/ioutil"
	"net/http"
)

type AzercellCabinetClient interface {
	GetToken() (string, error)
	SendSms(t string, s model.AzercellSmsBody) (*model.AzercellSmsResponse, error)
}

type AzercellCabinetClientImpl struct{}

func (c *AzercellCabinetClientImpl) GetToken() (string, error) {
	response := new(model.AzercellLoginResponse)
	endpoint := properties.Props.AzercellCabinetApiBaseUrl + "UserService/checkLoginDetailed/" +
		properties.Props.AzercellCabinetNumber + "/" + properties.Props.AzercellCabinetPassword
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, response)
	return response.LoginResponse.Token, err
}

func (c *AzercellCabinetClientImpl) SendSms(t string, s model.AzercellSmsBody) (*model.AzercellSmsResponse, error) {
	response := new(model.AzercellSmsResponse)
	endpoint := properties.Props.AzercellCabinetApiBaseUrl + "CustomerService/sendWebSmsNormal/" +
		properties.Props.AzercellCabinetNumber + "/" + t + "/" + properties.Props.AzercellCabinetNumber
	req, err := json.Marshal(s)
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(req))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, response)
	return response, err
}
