package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/anthonyLock/testify/assert"
)

func Test_main(t *testing.T) {
	// executeGameOfMatrix()
	executeGameOfWave()
	// executeGameOfKing()
	executeGameOfBlackAndWhite()
	executeGameOfLife()

	err := filepath.Walk("expected", func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)
		if !info.IsDir() {
			file1, err := ioutil.ReadFile(path)
			assert.NoError(t, err)
			file2Path := filepath.Join("output", info.Name())
			file2, err := ioutil.ReadFile(file2Path)
			assert.NoError(t, err)
			assert.Equal(t, file1, file2)
		}

		return nil
	})
	assert.NoError(t, err)
}
