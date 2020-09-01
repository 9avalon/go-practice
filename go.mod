module goworkspace

go 1.15

require (
	gee v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	xorm.io/cmd/xorm v0.0.0-20191108140657-006dbf24bb9b // indirect
	xorm.io/reverse v0.0.0-20200618084234-d29e5a0fd3ea // indirect
	xorm.io/xorm v1.0.3
)

replace gee => ./gee
