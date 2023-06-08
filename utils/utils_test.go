package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimUntilNum(t *testing.T) {
	s := TrimUntilNum("v18.16.0")
	assert.Equal(t, "18.16.0", s)
}

func TestGetFilenameFromUrl(t *testing.T) {
	s := GetFilenameFromUrl("https://nodejs.org/dist/v16.16.0/node-v16.16.0.tar.gz")
	assert.Equal(t, "node-v16.16.0.tar.gz", s)
}
