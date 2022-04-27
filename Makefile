userpb:
	protoc ./service/user_srv/userpb/user.proto --go_out=./service/user_srv/
	protoc ./service/user_srv/userpb/user.proto --go-grpc_out=./service/user_srv/

service:
	go run service/user_srv/main.go

api: 
	go run api/main.go

.PHONY: userpb service api

