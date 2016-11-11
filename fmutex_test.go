package fmutex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLockUnlock(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "fmutex")
	fm, err := New(filename)
	require.NoError(t, err)

	assert.Equal(t, true, fm.Lock())
	assert.Equal(t, true, fm.IsLocked())
	assert.Equal(t, nil, fm.Unlock())
	assert.Equal(t, false, fm.IsLocked())
}
