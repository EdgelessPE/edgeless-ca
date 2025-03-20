package translations

import "github.com/nicksnyder/go-i18n/v2/i18n"

var ZhMessages = []i18n.Message{
	{
		ID:    "unauthorized",
		Other: "未授权",
	},
	{
		ID:    "token_expired",
		Other: "令牌已过期",
	},
	{
		ID:    "invalid_token",
		Other: "无效的令牌",
	},
	{
		ID:    "no_token",
		Other: "未提供令牌",
	},
	{
		ID:    "email_sent",
		Other: "验证码已发送到您的邮箱",
	},
	{
		ID:    "email_error",
		Other: "发送验证码失败",
	},
	{
		ID:    "invalid_code",
		Other: "无效的验证码",
	},
	{
		ID:    "code_expired",
		Other: "验证码已过期",
	},
	{
		ID:    "code_error",
		Other: "验证码错误",
	},
	{
		ID:    "user_not_found",
		Other: "用户未找到",
	},
	{
		ID:    "password_error",
		Other: "密码错误",
	},
	{
		ID:    "one_hour_one_user",
		Other: "一小时只能发送一次验证码",
	},
	{
		ID:    "operation_too_frequent",
		Other: "操作过于频繁",
	},
	{
		ID:    "success",
		Other: "成功",
	},
	{
		ID:    "invalid_oauth_state",
		Other: "无效的 OAuth 状态",
	},
	{
		ID:    "token_exchange_error",
		Other: "令牌交换失败",
	},
	{
		ID:    "get_user_info_error",
		Other: "获取用户信息失败",
	},
	{
		ID:    "parse_user_info_error",
		Other: "解析用户信息失败",
	},
	{
		ID:    "get_user_email_error",
		Other: "获取用户邮箱失败",
	},
	{
		ID:    "parse_user_email_error",
		Other: "解析用户邮箱失败",
	},
}
