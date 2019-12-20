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
}

// Run migrations collection
func Run() {
	sort.Slice(Collection, func(i, j int) bool {
		return Collection[i].Timestamp() < Collection[j].Timestamp()
	})

	for _, c := range Collection {
		if !applied(c) {
			c.Apply()
		}
	}
}

func applied(mig Migration) bool {
	// get migrations from database
	migrations := []*migration{}

	for _, m := range migrations {
		if mig.Timestamp() == m.Timestamp {
			return true
		}
	}

	return false
}
