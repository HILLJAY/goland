package common

const (
	LoginMessageType = "LoginMes"
	LoginResType     = "LoginResMes"
	RegisterType     = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
}

type ResponseMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type RegisterMessage struct {
	User User `json:"user"`
}
