module github.com/l-e-t-i-c-i-a/microservices/shipping

go 1.24.0

replace github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping

//github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
require google.golang.org/grpc v1.78.0

require github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping v0.0.0-20260126115315-ddf2abd81ac5

require (
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260122232226-8e98ce8d340d // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
