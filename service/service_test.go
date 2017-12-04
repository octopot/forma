//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/service/mock_contract_test.go github.com/kamilsk/form-api/service DataLoader
package service_test
