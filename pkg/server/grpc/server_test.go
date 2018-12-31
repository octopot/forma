//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package grpc_test -destination mock_storage_test.go github.com/kamilsk/form-api/pkg/server/grpc ProtectedStorage
package grpc_test
