package ts_merge

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
	"github.com/stretchr/testify/assert"
)

type Traffic struct {
	Header string
	Value  string
}

type Serving struct {
	Name    string
	Traffic Traffic
}

type Serving2 struct {
	Name    string
	Traffic *Traffic
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
	mergo.Merge(originServing, defaultStruct, 
		mergo.WithOverride, mergo.WithTypeCheck)
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
	mergo.Merge(originServing, newStruct, 
		mergo.WithOverride, mergo.WithTypeCheck)
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

// 结构体中的指针部分也不会被新结构体的空指针所覆盖。
// tested on v0.3.6, v0.3.8
func TestMergeZeroPointer(t *testing.T) {
	originServing := &Serving2{
		Name: "orgin",
		Traffic: &Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	newStruct := &Serving2{
		Name: "newServing",
	}

	fmt.Println(originServing)
	fmt.Println(newStruct)
	assert.NotNil(t, originServing.Traffic, "Traffic should not be nil.")

	mergo.Merge(originServing, newStruct, mergo.WithOverride)

	fmt.Println(originServing)
	fmt.Println(newStruct)

	assert.NotNil(t, originServing.Traffic, "Traffic should not be nil.")
}

// v0.3.8: merge 时对于结构体中指针，指向不会变，会递归的 merge 指向的具体的值。
// v0.3.6: merge 会改变指针指向，为非零值时，不合理；为零值时，不会改变其指向；
func TestMergePointer(t *testing.T) {
	originServing := &Serving2{
		Name: "orgin",
		Traffic: &Traffic{
			Header: "origin_header",
			Value:  "origin_value",
		},
	}
	newStruct := &Serving2{
		Name: "newServing",
		Traffic: &Traffic{
			Header: "new_header",
			Value:  "new_value",
		},
	}

	fmt.Println(originServing)
	fmt.Println(newStruct)

	mergo.Merge(originServing, newStruct, mergo.WithOverride)

	fmt.Println(originServing)
	fmt.Println(originServing.Traffic)
	fmt.Println(newStruct)
}
