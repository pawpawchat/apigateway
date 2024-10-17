package dto

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
}

type RegisterOutput struct {
	DummyToken string `json:"dummy_token"`
}
