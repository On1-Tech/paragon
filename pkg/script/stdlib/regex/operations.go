package regex

import (
	"regexp"

	"github.com/kcarretto/paragon/pkg/script"
)

// Replace uses the golang regex lib to replace all occurences of the pattern in the old
// string into the new strong.
//
//go:generate go run ../gendoc.go -lib regex -func replace -param oldString@String -param pattern@String -param newString@String -retval replacedString@String -retval err@Error -doc "Replace uses the golang regex lib to replace all occurences of the pattern in the old string into the new strong."
//
// @callable:	regex.ReplaceString
// @param:		oldString			@String
// @param:		pattern				@String
// @param:		newString			@String
// @retval:		replacedString 		@String
// @retval:		err 				@Error
//
// @usage:		new_config = regex.replace(nginx_conf, "listen[\s]*80;", "listen 81;")
func Replace(oldString string, pattern string, newString string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllString(oldString, newString), nil
}

func replace(parser script.ArgParser) (script.Retval, error) {
	oldStr, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	pattern, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	newStr, err := parser.GetString(2)
	if err != nil {
		return nil, err
	}

	retVal, retErr := Replace(oldStr, pattern, newStr)
	return script.WithError(retVal, retErr), nil
}
