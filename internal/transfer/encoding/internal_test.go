package encoding

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

func TestOffers(t *testing.T) {
	for _, offer := range Offers() {
		assert.Contains(t, supported, offer)
	}
}

func TestHTMLEncoder_Encode(t *testing.T) {
	t.Run("marshal fails", func(t *testing.T) {
		enc := htmlEncoder{writer(nil)}
		assert.Error(t, enc.Encode(writer(nil)))
	})
	t.Run("data loss when recording", func(t *testing.T) {
		enc := htmlEncoder{writer(nil)}
		assert.Error(t, enc.Encode(domain.Schema{}))
	})
}

func TestYAMLEncoder_Encode(t *testing.T) {
	t.Run("marshal fails", func(t *testing.T) {
		enc := yamlEncoder{writer(nil), func(interface{}) ([]byte, error) { return nil, errors.New("problem marshal") }}
		assert.Error(t, enc.Encode(domain.Schema{}))
	})
	t.Run("data loss when recording", func(t *testing.T) {
		enc := yamlEncoder{writer(nil), func(interface{}) ([]byte, error) { return []byte("~"), nil }}
		assert.Error(t, enc.Encode(domain.Schema{}))
	})
}

type writer func(p []byte) (n int, err error)

func (writer) Write(p []byte) (n int, err error) {
	return len(p) - 1, nil
}

func (writer) MarshalHTML() ([]byte, error) {
	return nil, errors.New("marshal fails")
}
