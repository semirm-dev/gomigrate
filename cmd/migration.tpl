package migrations

import (
	"sort"
)

// Auto-generated file
// Feel free to edit

// Collection with all migrations
var Collection = []Migration{}

// Migration service
type Migration interface {
	Apply()
	Rollback()
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

			if err := saveMigrationHistory(c); err != nil {
				c.Rollback()
			}
		}
	}
}

func applied(mig Migration) bool {
	migrations := getMigrationsHistory()

	for _, m := range migrations {
		if mig.Timestamp() == m.Timestamp {
			return true
		}
	}

	return false
}

// To implement!!

func getMigrationsHistory() []*migration {
	// implement, get migrations from database
	return []*migration{}
}

func saveMigrationHistory(migration Migration) error {
	// implement, store migration in database
	return nil
}
