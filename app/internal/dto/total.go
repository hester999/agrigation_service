package dto

type TotalRequestDTO struct {
	ID          string `json:"id" example:"8d633c4c-ef75-475a-915a-ec5dd783dce9"` // Идентификатор пользователя
	ServiceName string `json:"service_name" example:"Yandex_plus"`                // Название услуги
	From        string `json:"from" example:"2025-06-01"`                         // Начальная дата диапазона (в формате YYYY-MM-DD)
	To          string `json:"to" example:"2025-07-31"`                           // Конечная дата диапазона (в формате YYYY-MM-DD)
}

type ResponseTotalDTO struct {
	TotalPrice int `json:"total_price" example:"1000"`
}
