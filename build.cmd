for /f "delims=" %%i in ('git describe --always') do set VERSION=%%i
set LDFLAGS="-X main.VERSION=%VERSION%"

set GOOS=windows
set GOARCH=amd64
go build -ldflags %LDFLAGS%