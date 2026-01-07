package response

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseWebData struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type AuthResponse struct {
	Code    int             `json:"code"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    ResponseWebData `json:"data,omitempty"`
}
