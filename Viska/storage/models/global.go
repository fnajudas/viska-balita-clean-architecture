package models

type Template2 struct {
	Items interface{} `json:"items"`
}

type Template1 struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Id          int    `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Template3 struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
