package ts_jsonpatch

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Patch []map[string]any

// will output '[]'
func TestEmptyJsonPatchOutput(t *testing.T) {
	p := Patch{}
	bytes, err := json.Marshal(p)
	assert.NoError(t, err)
	fmt.Println(string(bytes))
}
