//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package grpc_test -destination $PWD/pkg/server/grpc/mock_storage_test.go github.com/kamilsk/form-api/pkg/server/grpc ProtectedStorage
package grpc_test

import (
	_ "github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
)
