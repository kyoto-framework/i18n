package i18n

import (
	"reflect"
	"strings"
)

// TranslateDynamic extracts proper field from container according to provided language
func TranslateDynamic(lang string, container any, field string) string {
	// Behavior is different depending on type of container
	switch reflect.TypeOf(container).Kind() {
	case reflect.Map:
		// Cast to simplify code
		container := container.(map[string]any)
		// Check default (return without language suffix)
		if lang == Default {
			return container[field].(string)
		}
		// Check uppercase (camel case) suffix
		if val, exists := container[field+strings.ToUpper(lang)]; exists {
			return val.(string)
		}
		// Check lowercase (snake case) suffix
		if val, exists := container[field+"_"+strings.ToLower(lang)]; exists {
			return val.(string)
		}
		// Default
		return container[field].(string)
	case reflect.Struct:
		// Check default (return without language suffix)
		if lang == Default {
			return reflect.Indirect(reflect.ValueOf(container)).FieldByName(field).String()
		}
		// Check uppercase (camel case) suffix
		if val := reflect.Indirect(reflect.ValueOf(container)).FieldByName(field + strings.ToUpper(lang)).String(); len(val) != 0 && val != "<invalid Value>" {
			return val
		}
		// Default
		return reflect.Indirect(reflect.ValueOf(container)).FieldByName(field).String()
	default:
		panic("i18n: unknown container type for DynamicTranslate()")
	}
}
