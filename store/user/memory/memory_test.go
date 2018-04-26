package memory

import (
	"testing"

	usermodel "github.com/moqafi/harper/model/user"
)

func TestStore(t *testing.T) {
	store := New()
	var err error

	user1 := usermodel.User{ID: 1, Email: "user1@example.com"}
	user2 := usermodel.User{ID: 2, Email: "user2@example.com"}

	_ = store.Create(user1)
	_ = store.Create(user2)

	users, _ := store.List()

	if len(users) != 2 {
		t.Errorf("len(users) should be 2")
	}

	t.Logf("%v\n", users)

	err = store.Create(user1)
	if err == nil {
		t.Errorf("Error should be non nil")
	}

	t.Logf("%v\n", users)

	u1, err := store.Get(1)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	err = store.Delete(u1)
	users, _ = store.List()
	if len(users) != 1 {
		t.Errorf("Users should have a length of 1 after deleting")
	}

	t.Logf("%v\n", users)

	u2, _ := store.Get(2)
	u2.Email = "newuser2@example.com"

	err = store.Update(u2)
	if err != nil {
		t.Errorf("Error should be non-nil")
	}

	newU2, _ := store.Get(2)

	if newU2.Email != u2.Email {
		t.Errorf("User email not updated")
	}

	users, _ = store.List()
	t.Log(users)
}

func TestStoreCreate(t *testing.T) {

}

func TestStoreGet(t *testing.T) {

}

func TestStoreUpdate(t *testing.T) {

}

func TestStoreDelete(t *testing.T) {

}
