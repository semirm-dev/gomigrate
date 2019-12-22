package migrations

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/semirm-dev/go-dev/env"
	"github.com/sirupsen/logrus"
)

// Auto-generated file: https://github.com/semirm-dev/gomigrate
// Feel free to edit

type config struct {
	Dialect    string
	ConnString string
}

// Config for migrations
var Config = &config{
	Dialect: "postgres",
	ConnString: fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		env.Get("HOST", "localhost"),
		env.Get("PORT", "5432"),
		env.Get("DBNAME", "db_name"),
		env.Get("USER", "postgres"),
		env.Get("PASSWORD", "postgres"),
	),
}

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
	db, err := gorm.Open(Config.Dialect, Config.ConnString)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	sort.Slice(Collection, func(i, j int) bool {
		return Collection[i].Timestamp() < Collection[j].Timestamp()
	})

	migrations := getMigrationsHistory()

	for _, c := range Collection {
		if !applied(c, migrations) {
			c.Apply()

			if err := saveMigrationHistory(c); err != nil {
				c.Rollback()
			}
		}
	}
}

func applied(mig Migration, migrations []*migration) bool {
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
