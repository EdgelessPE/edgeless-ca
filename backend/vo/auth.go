package vo

type LoginPayload struct {
	Email   string `form:"email" binding:"required"`
	PwdHash string `form:"pwdHash" binding:"required"`
}
