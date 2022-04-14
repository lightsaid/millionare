userpb:
	protoc ./service/user_srv/userpb/user.proto --go_out=./service/user_srv/
	protoc ./service/user_srv/userpb/user.proto --go-grpc_out=./service/user_srv/


.PHONY: userpb

