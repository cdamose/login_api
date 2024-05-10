package dto

type UserProfile struct {
	UserId      string `json:"user_id"`
	IsVerified  bool   `json:"is_verified"`
	CreatedAt   string `json:"created_at"`
	VerfiedAt   string `json:"verfied_at"`
	PhoneNumber string `json:"phone_number"`
}

type VerifiedAccountResp struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	Message string `json:"message"`
}
