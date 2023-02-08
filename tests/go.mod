module gen/tests

go 1.16

require (
	gen v0.3.19
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	gorm.io/driver/mysql v1.4.6
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.5
	gorm.io/plugin/dbresolver v1.4.1
)

replace gen => ../
