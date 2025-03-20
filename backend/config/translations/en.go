package translations

import "github.com/nicksnyder/go-i18n/v2/i18n"

var EnMessages = []i18n.Message{
	{
		ID:    "unauthorized",
		Other: "Unauthorized",
	},
	{
		ID:    "token_expired",
		Other: "Token has expired",
	},
	{
		ID:    "invalid_token",
		Other: "Invalid token",
	},
	{
		ID:    "no_token",
		Other: "No token provided",
	},
	{
		ID:    "email_sent",
		Other: "Verification code has been sent to your email",
	},
	{
		ID:    "email_error",
		Other: "Failed to send verification code",
	},
	{
		ID:    "invalid_code",
		Other: "Invalid verification code",
	},
	{
		ID:    "code_expired",
		Other: "Verification code has expired",
	},
	{
		ID:    "code_error",
		Other: "Invalid verification code",
	},
	{
		ID:    "user_not_found",
		Other: "User not found",
	},
	{
		ID:    "password_error",
		Other: "Invalid password",
	},
	{
		ID:    "one_hour_one_user",
		Other: "One hour can only send one verification code",
	},
	{
		ID:    "operation_too_frequent",
		Other: "Operation too frequent",
	},
	{
		ID:    "success",
		Other: "Success",
	},
	{
		ID:    "invalid_oauth_state",
		Other: "Invalid OAuth state",
	},
	{
		ID:    "token_exchange_error",
		Other: "Token exchange failed",
	},
	{
		ID:    "get_user_info_error",
		Other: "Failed to get user info",
	},
	{
		ID:    "parse_user_info_error",
		Other: "Failed to parse user info",
	},
	{
		ID:    "get_user_email_error",
		Other: "Failed to get user email",
	},
	{
		ID:    "parse_user_email_error",
		Other: "Failed to parse user email",
	},
}
