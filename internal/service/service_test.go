//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package service_test -destination mock_storage_test.go go.octolab.org/ecosystem/forma/internal/service Storage
//go:generate mockgen -package service_test -destination mock_tracker_test.go go.octolab.org/ecosystem/forma/internal/service Tracker
package service_test
