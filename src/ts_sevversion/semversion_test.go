package ts_semversion

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

func TestSemversionRe(t *testing.T) {

	// version := "0.1.2-alpha+001"
	version := "1.12.0-b1-001"
	pattern := regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	r := pattern.FindStringSubmatch(version)

	m := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i == 0 {
			m["version"] = r[i]
		} else {
			m[name] = r[i]
		}
	}

	result, _ := json.MarshalIndent(m, "", "  ")
	fmt.Printf("%s\n", result)
}

/*
{
  "buildmetadata": "001",
  "major": "0",
  "minor": "1",
  "patch": "2",
  "prerelease": "alpha",
  "version": "0.1.2-alpha+001"
}
*/
