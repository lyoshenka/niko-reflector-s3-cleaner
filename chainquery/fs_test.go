package chainquery

import (
	"os"
	"testing"

	"reflector-s3-cleaner/shared"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndLoadSDHashes(t *testing.T) {
	existingHashes := []shared.StreamData{
		{"test1", 0, 0, true, false, false},
		{"test2", 0, 0, true, false, false},
		{"test3", 0, 0, true, false, false},
	}
	unresolvedHashes := []shared.StreamData{
		{"test4", 0, 0, false, false, false},
		{"test5", 0, 0, true, true, false},
		{"test6", 0, 0, true, false, true},
	}
	err := SaveHashes(existingHashes, "existing_sd_hashes.json")
	assert.NoError(t, err)
	defer os.Remove("existing_sd_hashes.json")
	err = SaveHashes(unresolvedHashes, "unresolved_sd_hashes.json")
	assert.NoError(t, err)
	defer os.Remove("unresolved_sd_hashes.json")

	loadedExistingHashes, err := LoadResolvedHashes("existing_sd_hashes.json")
	assert.NoError(t, err)
	assert.NotNil(t, loadedExistingHashes)
	assert.Len(t, loadedExistingHashes, 3)
	assert.ElementsMatch(t, existingHashes, loadedExistingHashes)

	loadedUnresolvedHashes, err := LoadResolvedHashes("unresolved_sd_hashes.json")
	assert.NoError(t, err)
	assert.NotNil(t, loadedUnresolvedHashes)
	assert.Len(t, loadedUnresolvedHashes, 3)
	assert.ElementsMatch(t, unresolvedHashes, loadedUnresolvedHashes)
}
