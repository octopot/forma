//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package grpc_test -destination mock_storage_test.go go.octolab.org/ecosystem/forma/internal/server/grpc ProtectedStorage
package grpc_test
