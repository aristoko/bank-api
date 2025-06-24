package model

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type UserResponseDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinedAt string `json:"joined_at"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToUserResponseDTO(u User) UserResponseDTO {
	return UserResponseDTO{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		JoinedAt: u.CreatedAt,
	}
}

func ToUserResponseDTOList(users []User) []UserResponseDTO {
	res := make([]UserResponseDTO, 0, len(users))
	for _, u := range users {
		res = append(res, ToUserResponseDTO(u))
	}
	return res
}
