package embedding

import "fmt"

type User struct {
	Name string
}

func (u User) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"name":%q}`, u.Name)), nil
}

type UserBio struct {
	Gender string
}

func (b UserBio) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"gender":%q}`, b.Gender)), nil
}
