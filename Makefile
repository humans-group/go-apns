gen:
	easyjson apns/api.go
	minimock -g -i ./apns.Client -o ./apns/ -s _mock.go

test:
	go test -v -race ./...