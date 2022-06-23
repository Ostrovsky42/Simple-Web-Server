package migrations

//go:generate go-bindata -pkg migrations -o bindata.go .

//migrate -source C:\Users\Serj\go\src\SimplyWebServer\migrations -database postgres://localhost:5432/transactionStore?sslmode=disable&user=postgres&password=gg" up
// migrate -path migrations -database "postgres://localhost:5432/transactionStore?sslmode=disable&user=postgres&password=gg" up

// migrate create -ext sql -dir migrations init
