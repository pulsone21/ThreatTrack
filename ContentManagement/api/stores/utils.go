package stores

import (
	"os"
	"path"
)

// Reads an sql file from the given path location
//
// Subpath starts without an ""
func LoadSQL(subpath string) (string, error) {
	path := path.Join("queries/", subpath)
	f, ioErr := os.ReadFile(path)
	if ioErr != nil {
		// handle error.
		return "", ioErr
	}
	sql := string(f)
	return sql, nil
}

type ScanFunc func(dest ...any) error
