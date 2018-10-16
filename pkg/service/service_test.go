//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_storage_test.go github.com/kamilsk/form-api/pkg/service Storage
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_tracker_test.go github.com/kamilsk/form-api/pkg/service Tracker
package service_test
