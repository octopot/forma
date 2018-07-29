package breaker_test

import (
	"os"
	"testing"
	"time"

	"github.com/kamilsk/go-kit/pkg/breaker"
)

func TestMultiplex(t *testing.T) {
	sleep := 100 * time.Millisecond

	start := time.Now()
	<-breaker.Multiplex(breaker.WithSignal(os.Interrupt), breaker.WithTimeout(sleep))
	end := time.Now()

	if expected, obtained := sleep, end.Sub(start); expected > obtained {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}

func TestMultiplex_WithoutChannels(t *testing.T) {
	<-breaker.Multiplex()
}

func TestWithDeadline(t *testing.T) {
	sleep := time.Now().Add(100 * time.Millisecond)

	<-breaker.WithDeadline(sleep)
	end := time.Now()

	if expected, obtained := sleep, end; expected.After(obtained) {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}

func TestWithSignal_NilSignal(t *testing.T) {
	<-breaker.WithSignal(nil)
}

func TestWithTimeout(t *testing.T) {
	sleep := 100 * time.Millisecond

	start := time.Now()
	<-breaker.WithTimeout(sleep)
	end := time.Now()

	if expected, obtained := sleep, end.Sub(start); expected > obtained {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}
