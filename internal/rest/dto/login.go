package dto

type LoginInput struct {
	Username string `json:"username" binding:"required"`
}

type LoginOutput struct {
	DummyToken string `json:"dummy_token"`
}
