module vprenner.com/comm

go 1.22.0

replace vprenner.com/protos => ../protos

require (
	github.com/gorilla/websocket v1.5.1
	google.golang.org/protobuf v1.32.0
	vprenner.com/protos v0.0.0-00010101000000-000000000000
)

require golang.org/x/net v0.17.0 // indirect
