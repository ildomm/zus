go clean -testcache
go test ./... -p 1 -run TestDatabase
go test ./... -p 1 -run "[^(TestDatabase)]"
