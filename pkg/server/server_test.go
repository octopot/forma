//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/pkg/server/mock_service_test.go github.com/kamilsk/form-api/pkg/server Service
package server_test
