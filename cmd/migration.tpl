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

type migration struct {
	Name      string
	Timestamp int64
	Applied   bool
}

// Run migrations collection
func Run() {
	sort.Slice(Collection, func(i, j int) bool {
		return Collection[i].Timestamp() < Collection[j].Timestamp()
	})

	// get migrations from database
	migrations := []*migration{}

	apply(migrations)
}

func apply(migrations []*migration) {
	for _, m := range migrations {
		if !m.Applied {
			for _, c := range Collection {
				if m.Timestamp == c.Timestamp() {
					c.Apply()
				}
			}
		}
	}
}
