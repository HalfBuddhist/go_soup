package ts_regex

import (
	"fmt"
	"regexp"
)

const ImageFullUrlPattern = `((?P<host>^[a-zA-Z0-9][-a-zA-Z0-9]*(\.[a-zA-Z0-9][-a-zA-Z0-9]*)+)/)?(?P<project>[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)/(?P<name>[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)(:(?P<version>[a-zA-Z0-9-_\.]*))?$`
const ImageShortUrlPattern = `(?P<project>^[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)/(?P<name>[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)(:(?P<version>[a-zA-Z0-9-_\.]*))?$`

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func GetParamsFromNamedPattern(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			if name != "" {
				paramsMap[name] = match[i]
			}
		}
	}
	return
}

func TS_regex_domain_name() {
	full_pattern := `^[a-zA-Z0-9][-a-zA-Z0-9]*(\.[a-zA-Z0-9][-a-zA-Z0-9]*)+$`
	reg := regexp.MustCompile(full_pattern)
	res := reg.Match([]byte("baid_u.com"))
	fmt.Println(res)
}

func TS_regex_image_url() {
	full_pattern := `((^[a-zA-Z0-9][-a-zA-Z0-9]*(\.[a-zA-Z0-9][-a-zA-Z0-9]*)+)/)?([a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)/([a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)(:([a-zA-Z0-9-_\.]*))?$`
	reg := regexp.MustCompile(full_pattern)
	res := reg.Match([]byte("-nvidia/l4t-tensorflow:r32.5.0-tf2.3-py3"))
	fmt.Println(res)
	short_pattern := `(^[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)/([a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)(:([a-zA-Z0-9-_\.]*))?$`
	res, _ = regexp.MatchString(short_pattern, "nvidia/l4t-tensorflow:r32.5.0-tf2.3-py3")
	fmt.Println(res)
}

func TS_Get_param() {
	res := GetParamsFromNamedPattern(ImageFullUrlPattern, "nvr.io/nvidia/l4t-tensorflow:r32.5.0-tf2.3-py3")
	fmt.Printf("%#v\n", res)
}
