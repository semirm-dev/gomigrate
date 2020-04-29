#### Usage

* install gomigrate tool
```sh
$ go install github.com/semirm-dev/gomigrate
```

* generate required templates from your project root
```
$ gomigrate tpl
```

* modify created config.yml

* create migration from your project root
```
$ gomigrate create -m=MyMigration1
```

* append migrations.Collection with previously created migration
```go
var Collection = []gomigrateCmd.MigrationDefinition{&MyMigration1{}}
```

* dont forget to execute migration command in main() and import database drivers
```go
package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/my-user/my-project/cmd"
)

func main() {
    cmd.Migration.Execute()
}
```

* when ready apply migrations from your project root
```sh
$ go run main.go migrate
```

Example of scaffolded code: https://github.com/semirm-dev/spotted-core/tree/master/migrations


#### TODO
- [ ] Automate manual appending to migrations.Collection
- [ ] Support to run specific migration (up+down)
