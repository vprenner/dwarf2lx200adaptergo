module vprenner.com/server

go 1.22.0

replace vprenner.com/comm => ../comm

replace vprenner.com/protos => ../protos

require vprenner.com/comm v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/websocket v1.5.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	vprenner.com/protos v0.0.0-00010101000000-000000000000 // indirect
)
