# BMI-GO
Go Lang Microservices demo

Pre-rquistes 
1. Protobuf (required for serializing structured data to JSON)- https://github.com/google/protobuf
2. Consul for service discovery

To Install and Use Protobuf:

Step 1: download https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-osx-x86_64.zip
Unzip to a suitable location example: /work/goApps/protoc-3.0.2-osx-x86_64/bin

Go to the ‘bin’ directory and check the version to make sure that you have the right thing, 
$ ./protoc --version

Step 2: Now add the path to protobuf compiler to an environment variable 
$	vi ~/.bash_profile
	Add an environment variable like this ‘export PROTOC=~/work/goApps/protoc-3.0.2-osx-x86_64/bin/./protoc’

Step 3: now set $GOPATH to location where you want to install protoc generator for Go, example export $GOPATH=~/work/goApps/

Step 4: Run `go get github.com/micro/protobuf/{proto,protoc-gen-go}
`. This will install this binary to $GOPATH/bin

Step 5: now add the above location to PATH environment variable like this
	$	vi ~/.bash_profile
	$ 	export PATH=$PATH:~/work/goApps/protobuf-master/protoc-gen-go


Step 6: open a new command window, to location of your .proto file
$ $PROTOC --go_out=plugins=micro:./pb namer.proto

Above command will compile your .proto file and generate a .go file in ‘pb’ directory.





