# HTML Entity Decoder
Simple Windows Golang UI to decode HTML Entities, e.g. `&quot;`

## build
env GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o ./htmldecoder.exe .