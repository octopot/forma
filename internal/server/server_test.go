//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package server_test -destination mock_service_test.go go.octolab.org/ecosystem/forma/internal/server Service
package server_test
