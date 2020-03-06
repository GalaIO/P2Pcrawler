package misc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type stringStruc struct {
	name string
}

func (s *stringStruc) String() string {
	return "stringStruc, name=" + s.name
}

type errorStruc struct {
	name string
}

func (s *errorStruc) Error() string {
	return "error, name=" + s.name
}

func Test2string(t *testing.T) {
	assert.Equal(t, "123", ToString("123"))
	assert.Equal(t, "123", ToString(123))
	assert.Equal(t, "1.2", ToString(1.2))
	assert.Equal(t, "[1,2,3]", ToString([]int{1, 2, 3}))
	assert.Equal(t, `{"name":"xiaoguo"}`, ToString(Dict{"name": "xiaoguo"}))
	assert.Equal(t, "stringStruc, name=xiaoguo", ToString(&stringStruc{name: "xiaoguo"}))
	assert.Equal(t, "error, name=xiaoguo", ToString(&errorStruc{name: "xiaoguo"}))
}
