#### Setup
* install gomigrate tool
```sh
$ go install github.com/semirm-dev/gomigrate
```

* generate required templates from your project root
```
$ gomigrate tpl -p=github.com/my-github-username/my-project
```

* modify created **_cmd/config.yml_**

* import database drivers in your main() and execute Migration cmd
```go
package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/my-github-username/my-project/cmd"
)

func main() {
    cmd.Migration.Execute()
}
```

### Usage
* create migration from your project root
```
$ gomigrate create -m=MyMigration1
```

* append migrations.Collection with previously created migration
```go
var Collection = []gomigrateCmd.MigrationDefinition{&MyMigration1{}}
```

* when ready apply migrations from your project root
```sh
$ go run main.go migrate
```

Example of scaffolded code: https://github.com/semirm-dev/spotted-core/tree/master/migrations


#### TODO
- [ ] Automate manual appending to migrations.Collection
- [ ] Support to run specific migration (up+down)
