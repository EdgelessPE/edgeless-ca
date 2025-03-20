package config

import (
	"log"

	"nep-keychain-backend/config/translations"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var enLocalizer *i18n.Localizer
var zhLocalizer *i18n.Localizer

func init() {
	bundle = i18n.NewBundle(language.SimplifiedChinese)
	loadTranslations()
	enLocalizer = i18n.NewLocalizer(bundle, language.English.String())
	zhLocalizer = i18n.NewLocalizer(bundle, language.SimplifiedChinese.String())
}

func loadTranslations() {
	// 加载英文翻译
	for _, message := range translations.EnMessages {
		bundle.AddMessages(language.English, &message)
	}

	// 加载中文翻译
	for _, message := range translations.ZhMessages {
		bundle.AddMessages(language.SimplifiedChinese, &message)
	}
}

// Translate 翻译消息
func Translate(messageID string, c *gin.Context) string {
	acceptLanguage := c.GetHeader("Accept-Language")
	localizer := enLocalizer
	if acceptLanguage == "zh" {
		localizer = zhLocalizer
	}

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	if err != nil {
		log.Printf("Error translating message: %v", err)
		return messageID
	}
	return message
}
