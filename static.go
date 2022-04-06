package i18n

// TranslateStatic takes static translation from provided path
func TranslateStatic(lang, page, group, key string) string {
	if len(files[page][lang][group][key]) == 0 {
		return files[page]["en"][group][key]
	}
	return files[page][lang][group][key]
}
