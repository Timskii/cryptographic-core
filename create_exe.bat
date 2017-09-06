::set GOARCH=386
::go build %GOPATH%\src\github.com\hyperledger\fabric\main.go
::mkdir %GOPATH%\src\github.com\hyperledger\fabric\exes\32
::move main.exe %GOPATH%\src\github.com\hyperledger\fabric\exes\32\main_32.exe
set GOARCH=amd64
go build %GOPATH%\src\github.com\hyperledger\fabric\main.go
mkdir %GOPATH%\src\github.com\hyperledger\fabric\exes\64
move main.exe %GOPATH%\src\github.com\hyperledger\fabric\exes\64\main_64.exe