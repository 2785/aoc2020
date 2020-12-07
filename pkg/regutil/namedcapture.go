package regutil

import (
	"fmt"
	"regexp"
)

// MustCaptureNamedGroup ..
func MustCaptureNamedGroup(re *regexp.Regexp, names []string, s string) (map[string]string, error) {
	match := re.FindStringSubmatch(s)
	result := make(map[string]string)
	for ind, name := range re.SubexpNames() {
		if ind != 0 && name != "" {
			result[name] = match[ind]
		}
	}

	for _, name := range names {
		val, ok := result[name]
		if !ok || val == "" {
			return nil, fmt.Errorf("%s not found in string %s", name, s)
		}
	}

	return result, nil
}
