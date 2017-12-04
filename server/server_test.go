//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/server/mock_contract_test.go github.com/kamilsk/form-api/server FormAPI,FormAPIService
package server_test
