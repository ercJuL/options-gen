package optionsgen_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	optionsgen "github.com/kazhuravlev/options-gen/options-gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	t.Run("simple_cases", func(t *testing.T) {
		t.Parallel()

		var testDirs []string
		{
			const testdataDir = "./testdata"
			fileInfos, err := os.ReadDir(testdataDir)
			require.NoError(t, err)

			for _, file := range fileInfos {
				if !file.IsDir() {
					continue
				}

				testDirs = append(testDirs, filepath.Join(testdataDir, file.Name()))
			}
		}

		for _, dir := range testDirs {
			dir := dir
			t.Run(dir, func(t *testing.T) {
				outFilename := filepath.Join(dir, "options_generated.go")
				expFilename := filepath.Join(dir, "options_generated.go.expected")
				paramsFilename := filepath.Join(dir, ".params.json")
				params := readParams(paramsFilename)

				err := optionsgen.Run(
					filepath.Join(dir, "options.go"),
					outFilename,
					"Options",
					"testcase",
					params.Defaults,
					true,
				)
				assert.NoError(t, err)

				helpEqualFiles(t, expFilename, outFilename)
			})
		}
	})

	t.Run("source_not_found", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		err := optionsgen.Run(
			filepath.Join(dir, "options.go"),
			filepath.Join(dir, "options_generated.go"),
			"Options",
			"testcase",
			optionsgen.Defaults{From: optionsgen.DefaultsFromTag, Param: ""},
			true,
		)
		assert.ErrorIs(t, err, syscall.ENOENT)
	})
}

type Params struct {
	Defaults optionsgen.Defaults `json:"defaults"`
}

func readParams(filename string) Params {
	params := Params{
		Defaults: optionsgen.Defaults{
			From:  optionsgen.DefaultsFromTag,
			Param: "",
		},
	}

	bb, err := os.ReadFile(filename)
	if err != nil {
		return params
	}

	if err := json.Unmarshal(bb, &params); err != nil {
		return params
	}

	return params
}

func helpEqualFiles(t *testing.T, filename1, filename2 string) {
	t.Helper()

	f1Bytes, err := os.ReadFile(filename1)
	require.NoError(t, err)

	f2Bytes, err := os.ReadFile(filename2)
	require.NoError(t, err)

	assert.Equal(t, string(f1Bytes), string(f2Bytes))
}
