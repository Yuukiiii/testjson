#!/bin/sh

#/usr/local/bin/protoc -I=./ --cpp_out=./ ./service_push.proto
/usr/local/bin/protoc -I=./ --cpp_out=./ ./service_common.proto
/usr/local/bin/protoc -I=./ --cpp_out=./ ./head.proto
/usr/local/bin/protoc -I=./ --cpp_out=./ ./service_status.proto
/usr/local/bin/protoc -I=./ --cpp_out=./ ./tcpd_service.proto
/usr/local/bin/protoc -I=./ --cpp_out=./ ./control_exchange.proto

mv -f *.pb.h ../src/ 
mv -f *.pb.cc ../src/ 
#/usr/local/bin/protoc -I=./ ./head.proto -o ./head.pb
#/usr/local/bin/protoc -I=./ ./control_exchange.proto -o ./control_exchange.pb
#/usr/local/bin/protoc -I=./ ./service_common.proto -o ./service_common.pb
#/usr/local/bin/protoc -I=./ ./service_push.proto -o ./service_push.pb
