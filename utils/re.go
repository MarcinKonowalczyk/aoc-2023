package utils

import "regexp"

// https://stackoverflow.com/a/39635221
func GetParamsCompiledRe(re *regexp.Regexp, s string) map[string]string {
	match := re.FindStringSubmatch(s)
	params := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[name] = match[i]
		}
	}
	return params
}

func GetParamsStringRe(res string, s string) (map[string]string, error) {
	re, err := regexp.Compile(res)
	if err != nil {
		return nil, err
	}
	return GetParamsCompiledRe(re, s), nil
}
