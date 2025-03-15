package vo

type LoginPayload struct {
	Email   string `form:"email" binding:"required"`
	PwdHash string `form:"pwdHash" binding:"required"`
}

type RecoverPayload struct {
	Email   string `form:"email" binding:"required"`
	Code    string `form:"code" binding:"required"`
	PwdHash string `form:"pwdHash" binding:"required"`
}

type SendVerifyCodePayload struct {
	Email string `form:"email" binding:"required"`
}
