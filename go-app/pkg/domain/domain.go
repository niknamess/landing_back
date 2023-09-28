package domain

import "time"

type RegisterRequest struct {
	Email    string `json:"email"`
	Fio      *Fio   `json:"fio"`
	Password string `json:"password"`
	Driver   bool   `json:"driver"`
	Customer bool   `json:"customer"`
}

type RegisterCheckRequest struct {
	Email string `json:"email"`
}

type RegisterCheckInputRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Fio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
type RegisterResponse struct {
	Status string `json:"status"`
}

type FormToAction struct {
	UserID      string    `json:"user_id"`
	WhereFrom   string    `json:"where_from"`
	WhereTo     string    `json:"where_to"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

type FormsResponse struct {
	Forms []*FormToAction `json:"forms"`
}

type GetFormsRequest struct {
	UserID string `json:"user_id"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
