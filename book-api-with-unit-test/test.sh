export DB_USERNAME_TEST=tegarap
export DB_PASSWORD_TEST=t00r!Roo
export DB_HOST_TEST=localhost
export DB_PORT_TEST=3306
export DB_NAME_TEST=tegar_store
export SERV_PORT_TEST=9090
export SECRET_JWT=legal

go test -v -coverprofile=coverage.out ./controllers/...
go tool cover -func coverage.out