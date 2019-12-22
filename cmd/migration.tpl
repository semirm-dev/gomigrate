package migrations

import (
	"fmt"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/semirm-dev/go-dev/env"
	"github.com/sirupsen/logrus"
)

// Auto-generated file: https://github.com/semirm-dev/gomigrate
// Feel free to edit

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

type config struct {
	Dialect    string
	ConnString string
}

// Migration service
type Migration interface {
	Name() string
	Apply()
	Rollback()
	Timestamp() int64
}

type migration struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Timestamp int64
	CreatedAt time.Time
}

// Run migrations collection
func Run() {
	db, err := gorm.Open(Config.Dialect, Config.ConnString)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&migration{})

	sort.Slice(Collection, func(i, j int) bool {
		return Collection[i].Timestamp() < Collection[j].Timestamp()
	})

	migrations := getMigrationsHistory(db)

	for _, c := range Collection {
		if !applied(c, migrations) {
			c.Apply()

			if err := saveMigrationHistory(c, db); err != nil {
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

func getMigrationsHistory(db *gorm.DB) []*migration {
	var migrations = []*migration{}

	db.Find(&migrations)

	return migrations
}

func saveMigrationHistory(m Migration, db *gorm.DB) error {
	return db.Create(&migration{
		Name:      m.Name(),
		Timestamp: m.Timestamp(),
	}).Error
}
