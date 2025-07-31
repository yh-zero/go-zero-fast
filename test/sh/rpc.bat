@REM 项目根目录  运行例子1： .\test\sh\rpc.bat sys sys

set urlName=%1
set rpcName=%2

goctl rpc protoc service/%urlName%/rpc/desc/%rpcName%.proto --go_out=service/%urlName%/rpc/ --go-grpc_out=service/%urlName%/rpc/ --zrpc_out=service/%urlName%/rpc/ -m --style=go_zero --home=./template/185/

@REM goctl rpc protoc application\usre\rpc\usre.proto --go_out=application\usre\rpc --go-grpc_out=application\usre\rpc --zrpc_out=application\usre\rpc\


@REM goctl rpc protoc ./user.proto --go_out=. --go-grpc_out=. --zrpc_out=./
