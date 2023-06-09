package archiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompress(t *testing.T) {
	err := Decompress("../../test/node-v16.16.0.tar.gz", "../../test")
	assert.Equal(t, nil, err)
}
