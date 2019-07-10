thrift -r --gen go shared.thrift
thrift -r --gen go tutorial.thrift
cd gen-go/shared && go mod init shared && cd -
cd gen-go/tutorial && go mod init tutorial && cd -
go build -o main
./main -server=true