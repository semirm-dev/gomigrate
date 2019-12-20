#### Migrations scaffolding

#### Usage

* install gomigrate tool
```sh
$ go install github.com/semirm-dev/gomigrate
```

* generate required templates 
```
$ gomigrate tpl
```

* create migration from your project root
```
$ gomigrate --create=MyMigration1
```

* append migrations.Collection with previously created migration
```go
var Collection = []Migration{&MyMigration1{}}
```
