package ts_merge

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
)

type Traffic struct {
	Header string
	Value  string
}

type Serving struct {
	Name    string
	Traffic Traffic
}

// 空值不会进行赋值
func TestMergeDefaultValue(t *testing.T) {
	originServing := &Serving{
		Name: "orgin",
		Traffic: Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	defaultStruct := &Serving{}
	fmt.Println(originServing)
	fmt.Println(defaultStruct)

	// WithOverride
	mergo.Merge(originServing, defaultStruct, mergo.WithOverride)
	fmt.Println(originServing)
	fmt.Println(defaultStruct)

	// WithTypeCheck
	mergo.Merge(originServing, defaultStruct, mergo.WithOverride, mergo.WithTypeCheck)
	fmt.Println(originServing)
	fmt.Println(defaultStruct)

	// MergeWithOverwrite
	mergo.MergeWithOverwrite(originServing, defaultStruct)
	fmt.Println(originServing)
	fmt.Println(defaultStruct)

	// MapWithOverwrite
	mergo.MapWithOverwrite(originServing, defaultStruct)
	fmt.Println(originServing)
	fmt.Println(defaultStruct)
}

// 嵌套里面的空值也不会进行赋值, 只赋值嵌套中的非零值部分。
func TestMergeDefaultValueNest(t *testing.T) {
	originServing := &Serving{
		Name: "orgin",
		Traffic: Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	newStruct := &Serving{
		Traffic: Traffic{
			Value: "xxx",
		},
	}
	fmt.Println(originServing)
	fmt.Println(newStruct)

	// WithOverride
	mergo.Merge(originServing, newStruct, mergo.WithOverride)
	fmt.Println(originServing)
	fmt.Println(newStruct)

	// WithTypeCheck
	originServing = &Serving{
		Name: "orgin",
		Traffic: Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	mergo.Merge(originServing, newStruct, mergo.WithOverride, mergo.WithTypeCheck)
	fmt.Println(originServing)
	fmt.Println(newStruct)

	// MergeWithOverwrite
	originServing = &Serving{
		Name: "orgin",
		Traffic: Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	mergo.MergeWithOverwrite(originServing, newStruct)
	fmt.Println(originServing)
	fmt.Println(newStruct)

	// MapWithOverwrite
	originServing = &Serving{
		Name: "orgin",
		Traffic: Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	mergo.MapWithOverwrite(originServing, newStruct)
	fmt.Println(originServing)
	fmt.Println(newStruct)
}
