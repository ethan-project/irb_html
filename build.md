removing unused dependencies
------------- 
go mod tidy

listing dependencies
-------------
go list -m all

checks that dependencies
-------------
go mod verify


build
=============
/root/go/swag init -d ./ipsap/ -o ./ipsap/docs -g ipsap.go
go build ipsap.go

run
=============
go run ./ipsap_main.go '/root/go/config/*.toml' &