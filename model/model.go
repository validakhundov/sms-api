package model

type Sms struct {
	Number string `json:"number"`
	Text   string `json:"text"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type AzercellSmsResponse struct {
	StringResponse AzercellSmsStringResponse `json:"stringResponse"`
}

type AzercellSmsStringResponse struct {
	ResponseMessage string `json:"responseMessage"`
}

type AzercellSmsBody struct {
	Date    string `json:"date"`
	Message string `json:"message"`
	SendTo  string `json:"sendto"`
}

type AzercellLoginResponse struct {
	LoginResponse DetailedLoginResponse `json:"DetailedLoginResponse"`
}

type DetailedLoginResponse struct {
	Token string `json:"token"`
}
