package main

import (
	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

// Translate English To Japanese
func engToJapMsg(msgContent string) string {
	translatedText, err := gtranslate.TranslateWithParams(
		msgContent,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | ja: %s \n", msgContent, translated)
}

// Translate English To Spainish
func engToSpanishMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.Spanish)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | spainish: %s \n", msgContent, translatedText)
}

// Translate English To Chinese
func engToChineseMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.SimplifiedChinese)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | simplified chinese: %s \n", msgContent, translatedText)
}

// Translate English To German
func engToGermanMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.German)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | german: %s \n", msgContent, translatedText)
}

// Translate Japanese To English
func japToEngMsg(msgContent string) string {
	translatedText, err := gtranslate.TranslateWithParams(
		msgContent,
		gtranslate.TranslationParams{
			From: "ja",
			To:   "en",
		},
	)
	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | ja: %s \n", msgContent, translated)
}

// Translate Spainish To English
func spanishToEngMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.Spanish, language.English)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | spainish: %s \n", msgContent, translatedText)
}

// Translate Chinese To English
func chineseToEngMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.SimplifiedChinese, language.English)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | simplified chinese: %s \n", msgContent, translatedText)
}

// Translate German To English
func germanToEngMsg(msgContent string) string {
	translatedText, err := gtranslate.Translate(msgContent, language.German, language.English)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | german: %s \n", msgContent, translatedText)
}

//test whatsapp
