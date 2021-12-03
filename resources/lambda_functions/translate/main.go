package main

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

type TranslateEvent struct {
	Text       string   `json:"sourceText"`
	SourceLang LangCode `json:"sourceLang"`
	TargetLang LangCode `json:"targetLang"`
}

type LangCode string

func (l LangCode) normalize(defaultCode string) string {
	lang := strings.ToLower(string(l))
	if _, ok := availableLangCodes[lang]; ok {
		return lang
	}

	return defaultCode
}

func handleRequest(ctx context.Context, event TranslateEvent) (string, error) {
	c, err := newTranslateClient()
	if err != nil {
		return "", err
	}

	sourceLang := event.SourceLang.normalize("ja")
	targetLang := event.TargetLang.normalize("en")

	res, err := translateText(c, event.Text, sourceLang, targetLang)
	if err != nil {
		return "", err
	}

	return res, nil
}

func newTranslateClient() (*translate.Client, error) {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		return nil, errors.New("Environment variable AWS_DEFAULT_REGION is undefined.")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, errors.New("Failed to laod AWS deault config.")
	}

	client := translate.NewFromConfig(cfg)

	return client, nil
}

func translateText(c *translate.Client, text, sourceLang, targetLang string) (string, error) {
	in := &translate.TranslateTextInput{
		SourceLanguageCode: aws.String(sourceLang),
		TargetLanguageCode: aws.String(targetLang),
		Text:               aws.String(text),
	}

	out, err := c.TranslateText(context.Background(), in)
	if err != nil {
		return "", err
	}

	translated := aws.ToString(out.TranslatedText)

	return translated, nil
}

func main() {
	lambda.Start(handleRequest)
}

var availableLangCodes map[string]string = map[string]string{
	"af":    "Afrikaans",
	"am":    "Amharic",
	"ar":    "Arabic",
	"az":    "Azerbaijani",
	"bg":    "Bulgarian",
	"bn":    "Bengali",
	"bs":    "Bosnian",
	"ca":    "Catalan",
	"cs":    "Czech",
	"cy":    "Welsh",
	"da":    "Danish",
	"de":    "German",
	"el":    "Greek",
	"en":    "English",
	"es-MX": "Spanish (Mexico)",
	"es":    "Spanish",
	"et":    "Estonian",
	"fa-AF": "Dari",
	"fa":    "Farsi (Persian)",
	"fi":    "Finnish",
	"fr-CA": "French (Canada)",
	"fr":    "French",
	"ga":    "Irish",
	"gu":    "Gujarati",
	"ha":    "Hausa",
	"he":    "Hebrew",
	"hi":    "Hindi",
	"hr":    "Croatian",
	"ht":    "Haitian Creole",
	"hu":    "Hungarian",
	"hy":    "Armenian",
	"id":    "Indonesian",
	"is":    "Icelandic",
	"it":    "Italian",
	"ja":    "Japanese",
	"ka":    "Georgian",
	"kk":    "Kazakh",
	"kn":    "Kannada",
	"ko":    "Korean",
	"lt":    "Lithuanian",
	"lv":    "Latvian",
	"mk":    "Macedonian",
	"ml":    "Malayalam",
	"mn":    "Mongolian",
	"mr":    "Marathi",
	"ms":    "Malay",
	"mt":    "Maltese",
	"nl":    "Dutch",
	"no":    "Norwegian",
	"pa":    "Punjabi",
	"pl":    "Polish",
	"ps":    "Pashto",
	"pt-PT": "Portuguese (Portugal)",
	"pt":    "Portuguese",
	"ro":    "Romanian",
	"ru":    "Russian",
	"si":    "Sinhala",
	"sk":    "Slovak",
	"sl":    "Slovenian",
	"so":    "Somali",
	"sq":    "Albanian",
	"sr":    "Serbian",
	"sv":    "Swedish",
	"sw":    "Swahili",
	"ta":    "Tamil",
	"te":    "Telugu",
	"th":    "Thai",
	"tl":    "Filipino, Tagalog",
	"tr":    "Turkish",
	"uk":    "Ukrainian",
	"ur":    "Urdu",
	"uz":    "Uzbek",
	"vi":    "Vietnamese",
	"zh-TW": "Chinese (Traditional)",
	"zh":    "Chinese (Simplified)",
}
