package dto

type ResponseJson struct {
	Success string    `json:"success"`
	Data    []Holiday `json:"data"`
}
