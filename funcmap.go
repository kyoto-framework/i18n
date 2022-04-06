package i18n

import "html/template"

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"translateStatic":  TranslateStatic,
		"translateDynamic": TranslateDynamic,
	}
}
