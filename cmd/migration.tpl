package migrations

import (
	"sort"
)

// Collection with all migrations
var Collection = []Migration{}

// Migration service
type Migration interface {
	Apply()
	Timestamp() int64
}

// Run migrations collection
func Run() {
	sort.Slice(Collection, func(i, j int) bool {
		return Collection[i].Timestamp() < Collection[j].Timestamp()
	})

	for _, m := range Collection {
		m.Apply()
	}
}
