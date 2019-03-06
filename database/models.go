package database

import (
	"strconv"

	"cloud.google.com/go/datastore"
)

type User struct {
	Key      *datastore.Key `datastore:"__key__"`
	Username string         `datastore:"username"`
	Email    string         `datastore:"email"`

	// If due to historical reasons, there are two types in Datastore: string and int.
	// Can not get User Entity by using:
	// query := datastore.NewQuery("User")
	// _, err := client.GetAll(ctx, query, &users)
	// It will cause ErrFieldMismatch error.
	Age int64 `datastore:"age"`
}

// Customizing datastore.Keyloader to slove ErrFieldMismatch.
func (user *User) Load(properties []datastore.Property) error {
	for i, p := range properties {
		if p.Name == "age" {
			switch p.Value.(type) {
			case string:
				a := p.Value.(string)
				aInt, _ := strconv.Atoi(a)
				properties[i].Value = int64(aInt)
			}
		}
	}

	return datastore.LoadStruct(user, properties)
}

func (user *User) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(user)
}

func (user *User) LoadKey(k *datastore.Key) error {
	user.Key = k
	return nil
}
