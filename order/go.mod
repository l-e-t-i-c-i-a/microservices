module github.com/l-e-t-i-c-i-a/microservices/order

go 1.24.0

require (
	//github.com/l-e-t-i-c-i-a/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.78.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260122232226-8e98ce8d340d // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/l-e-t-i-c-i-a/microservices-proto/golang/order => ../../microservices-proto/golang/order

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/l-e-t-i-c-i-a/microservices-proto/golang/order v0.0.0-20260126115315-ddf2abd81ac5
	github.com/l-e-t-i-c-i-a/microservices-proto/golang/payment v0.0.0-20260126115315-ddf2abd81ac5
	github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping v0.0.0-20260126115315-ddf2abd81ac5
//github.com/l-e-t-i-c-i-a/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
//github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
)

replace github.com/l-e-t-i-c-i-a/microservices-proto/golang/payment => ../../microservices-proto/golang/payment

replace github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
