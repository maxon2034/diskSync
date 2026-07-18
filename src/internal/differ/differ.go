package differ

import "diskSync/src/internal/storage"

func Diff(source, dest []storage.File) []storage.File {
	destSet := make(map[string]struct{}, len(dest))

	for _, f := range dest {
		destSet[f.Name] = struct{}{}
	}

	var diff []storage.File
	for _, f := range source {
		if _, exists := destSet[f.Name]; !exists {
			diff = append(diff, f)
		}
	}
	return diff
}
