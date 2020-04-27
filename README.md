#### Usage

* install gomigrate tool
```sh
$ go install github.com/semirm-dev/gomigrate
```

* generate required templates from your project root
```
$ gomigrate tpl -p=my-github-username/my-project
```

* create migration from your project root
```
$ gomigrate create -m=MyMigration1
```

* append migrations.Collection with previously created migration
```go
var Collection = []Migration{&MyMigration1{}}
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


#### TODO
- [ ] Automate required appending to migrations.Collection
- [ ] Support to run specific migration (up+down)
