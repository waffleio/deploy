package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOutput(t *testing.T) {
	output := getOutput()
	assert.Equal(t, "Done.", output)
}
