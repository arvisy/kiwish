# ------------------------------------------------------------------ #
#                              protobuf                              #
# ------------------------------------------------------------------ #
.PHONY:pb/gen
pb/gen:
	@protoc --proto_path=. \
		--go_out=. --go_opt=module=ms-cart \
		--go-grpc_out=. --go-grpc_opt=module=ms-cart \
		./proto/*.proto

.PHONY:pb/rm
pb/rm:
	@rm -rf pb/*