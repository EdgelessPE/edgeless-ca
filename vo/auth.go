package vo

type LoginPayload struct {
	Name    string `form:"name" binding:"required"`
	PwdHash string `form:"pwdHash" binding:"required"`
}

type RegisterPayload struct {
	Email   string `form:"email" binding:"required"`
	Name    string `form:"name" binding:"required"`
	PwdHash string `form:"pwdHash" binding:"required"`
}
