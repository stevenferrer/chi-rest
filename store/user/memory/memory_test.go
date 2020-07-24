package memory

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	usermodel "github.com/sf9v/harper/model/user"
)

func TestStoreList(t *testing.T) {
	store := New()
	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1user1user1"),
	}
	user2 := usermodel.User{
		Email:    "user2@example.com",
		Password: []byte("user2user2user2"),
	}

	user1, _ = store.Create(user1)
	user2, _ = store.Create(user2)

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
		Password: []byte("user1user1"),
	}
	user2 := usermodel.User{
		Email:    "user2@example.com",
		Password: []byte("user2user2user2"),
	}

	user1, _ = store.Create(user1)
	user2, _ = store.Create(user2)

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

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1user1user1"),
	}
	user1, _ = store.Create(user1)

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

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1user1user1"),
	}
	user1, _ = store.Create(user1)

	newPwd := "newuser1password"
	user1.Password = []byte(newPwd)

	newUser1, err := store.UpdateByEmail(user1.Email, user1)
	if err != nil {
		t.Fatal(err)
	}

	if err := bcrypt.CompareHashAndPassword(newUser1.Password, []byte(newPwd)); err != nil {
		t.Fatal("password was not updated", err)
	}
}

func TestStoreDelete(t *testing.T) {
	store := New()

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1user1user1"),
	}
	user1, _ = store.Create(user1)

	user1, err := store.Delete(user1)
	if err != nil {
		t.Error(err)
	}

	_, err = store.GetByID(user1.ID)
	if err == nil {
		t.Error("Expecting error because user1 is not in the store")
	}
}

func TestValidate(t *testing.T) {
	store := New()
	var err error

	user1 := usermodel.User{
		Email:    "user1example.com",
		Password: []byte("user1"),
	}
	user1, err = store.Create(user1)
	if err == nil {
		t.Errorf("expecting non-error invalid email got: %v\n", err)
	}

	t.Logf("validation ok: %v", err)
}
