gen:
	easyjson notification.go api.go
	minimock -g -i Client -o ./ -s _mock.go