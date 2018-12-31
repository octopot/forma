//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package server_test -destination mock_service_test.go github.com/kamilsk/form-api/pkg/server Service
package server_test
