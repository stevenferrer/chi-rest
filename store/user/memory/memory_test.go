package memory

import (
	"log"
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

	log.Println(users)

	err = store.Create(user1)
	if err == nil {
		t.Errorf("Error should be non nil")
	}

	log.Println(err)

	u1, err := store.Get(1)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	log.Println(u1)
}
