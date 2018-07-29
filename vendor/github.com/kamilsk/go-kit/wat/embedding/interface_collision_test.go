package embedding_test

import (
	"encoding/json"
	"testing"

	"github.com/kamilsk/go-kit/wat/embedding"
)

func TestInterfaceCollision_Marshaler(t *testing.T) {
	// check that Marshaler interface works properly for User
	j, _ := json.Marshal(embedding.User{Name: "John"})
	if expected := `{"name":"John"}`; string(j) != expected {
		t.Errorf("obtained: %q; expected: %q", j, expected)
	}

	// check that Marshaler interface works properly for UserBio
	j, _ = json.Marshal(embedding.UserBio{Gender: "male"})
	if expected := `{"gender":"male"}`; string(j) != expected {
		t.Errorf("obtained: %q; expected: %q", j, expected)
	}

	// but what will be if we embedding struct which implements Marshaler?
	j, _ = json.Marshal(struct {
		embedding.User
		Age uint
	}{User: embedding.User{Name: "John"}, Age: 30})
	if expected := `{"name":"John"}`; string(j) != expected { // wat: where is the Age?
		t.Errorf("obtained: %q; expected: %q", j, expected)
	}

	// and what will happen if we combine them?
	j, _ = json.Marshal(struct {
		embedding.User
		embedding.UserBio
	}{User: embedding.User{Name: "John"}, UserBio: embedding.UserBio{Gender: "male"}})
	if expected := `{"Name":"John","Gender":"male"}`; string(j) != expected { // wat: why not "name" and "gender"?
		t.Errorf("obtained: %q; expected: %q", j, expected)
	}
}
