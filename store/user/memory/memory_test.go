package memory

import (
	"testing"

	usermodel "github.com/moqafi/harper/model/user"
)

func TestStoreList(t *testing.T) {
	store := New()
	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: "user1",
	}
	user2 := usermodel.User{
		Email:    "user2@example.com",
		Password: "user2",
	}

	_ = store.Create(user1)
	_ = store.Create(user2)

	users, err := store.List()
	if err != nil {
		t.Error(err)
	}

	if len(users) != 2 {
		t.Fatal("User len expected to be 2")
	}
}

func TestStoreCreate(t *testing.T) {
	store := New()
	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: "user1",
	}
	user2 := usermodel.User{
		Email:    "user2@example.com",
		Password: "user2",
	}

	_ = store.Create(user1)
	_ = store.Create(user2)

	users, err := store.List()
	if err != nil {
		t.Error(err)
	}

	if len(users) != 2 {
		t.Fatal("User len expected to be 2")
	}
}

func TestStoreGet(t *testing.T) {
	store := New()
	var err error

	user1 := usermodel.User{Email: "user1@example.com", Password: "user1"}
	_ = store.Create(user1)

	sameUser1, err := store.GetByEmail(user1.Email)
	if err != nil {
		t.Error(err)
	}

	if user1.Email != sameUser1.Email {
		t.Fatal("user1 should have the same email with sameUser1")
	}
}

func TestStoreUpdate(t *testing.T) {
	store := New()

	user1 := usermodel.User{Email: "user1@example.com", Password: "user1"}
	_ = store.Create(user1)

	user1.Password = "newuser1password"

	err := store.UpdateByEmail(user1.Email, user1)

	newUser1, err := store.GetByEmail(user1.Email)
	if err != nil {
		t.Error(err)
	}

	if user1.Password != newUser1.Password {
		t.Fatal("user should have the same password with newUser1")
	}
}

func TestStoreDelete(t *testing.T) {
	store := New()

	user1 := usermodel.User{Email: "user1@example.com", Password: "user1"}
	_ = store.Create(user1)

	err := store.Delete(user1)
	if err == nil {
		t.Error(err)
	}

	_, err = store.GetByID(user1.ID)
	if err == nil {
		t.Error("Expecting error because user1 is not in the store")
	}
}
