package dto

type ErrDTO400 struct {
	Message string `json:"message" example:"Bad Request"`
	Code    int    `json:"code" example:"400"`
}

type ErrDTO404 struct {
	Message string `json:"message" example:"Not Found"`
	Code    int    `json:"code" example:"404"`
}

type ErrDTO500 struct {
	Message string `json:"message" example:"internal server error"`
	Code    int    `json:"code" example:"500"`
}

type ErrDTOArr struct {
	Message string      `json:"message" example:"Not Found"`
	Code    int         `json:"code" example:"404"`
	Data    interface{} `json:"data"`
}
