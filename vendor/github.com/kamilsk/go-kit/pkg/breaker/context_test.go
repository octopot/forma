// +build go1.7

package breaker_test

import (
	"context"
	"testing"
	"time"

	"github.com/kamilsk/go-kit/pkg/breaker"
)

func TestWithContext(t *testing.T) {
	sleep := 100 * time.Millisecond
	ctx := breaker.WithContext(context.TODO(), breaker.WithTimeout(sleep))

	start := time.Now()
	<-ctx.Done()
	end := time.Now()

	if expected, obtained := sleep, end.Sub(start); expected > obtained {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}
