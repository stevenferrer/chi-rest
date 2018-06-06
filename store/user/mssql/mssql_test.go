package mssql

import (
	//	"log"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	usermodel "github.com/moqafi/harper/model/user"
)

func getDb() (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;database=%s;",
		"localhost\\mssql2016express",
		"sa",
		"sa",
		"harper",
	)

	db, err := gorm.Open("mssql", connStr)
	if err != nil {
		return nil, err
	}

	//	var (
	//		sqlVersion string
	//	)
	//	rows, err := db.DB().Query("select @@version")
	//	if err != nil {
	//		return nil, err
	//	}

	//	for rows.Next() {
	//		err := rows.Scan(&sqlVersion)
	//		if err != nil {
	//			return nil, err
	//		}
	//		log.Println(sqlVersion)
	//	}

	return db, nil
}

func TestCreate(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Fatal(err)
	}
	// migrate
	err = db.AutoMigrate(new(usermodel.User)).Error
	if err != nil {
		t.Fatal(err)
	}

	store := New(db)

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1"),
	}

	user1, err = store.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	// delete test data
	err = db.Unscoped().Delete(&user1).Error
	if err != nil {
		t.Fatal(err)
	}
}

func TestList(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Fatal(err)
	}
	// migrate
	err = db.AutoMigrate(new(usermodel.User)).Error
	if err != nil {
		t.Fatal(err)
	}

	store := New(db)

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1"),
	}

	user1, err = store.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	users, err := store.List()
	if err != nil {
		t.Fatal(err)
	}

	for _, u := range users {
		t.Logf("ID: %d, Email: %s\n", u.ID, u.Email)
	}

	// delete test data
	err = db.Unscoped().Delete(&user1).Error
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByEmail(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Fatal(err)
	}
	// migrate
	err = db.AutoMigrate(new(usermodel.User)).Error
	if err != nil {
		t.Fatal(err)
	}

	store := New(db)

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1"),
	}

	sameUser1, err := store.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: update this, get is missing here

	if sameUser1.Email != user1.Email {
		t.Fatalf("expected same user email got different: %s != %s", user1.Email, sameUser1.Email)
	}

	// delete test data
	err = db.Unscoped().Delete(&user1).Error
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByID(t *testing.T) {

}

func TestUpdateByEmail(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Fatal(err)
	}
	// migrate
	err = db.AutoMigrate(new(usermodel.User)).Error
	if err != nil {
		t.Fatal(err)
	}

	store := New(db)

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1"),
	}

	sameUser1, err := store.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	// save old email
	oldEmail := sameUser1.Email
	//change email
	sameUser1.Email = "newuser1@example.com"
	sameUser1.Password = []byte("newuser1")

	newUser1, err := store.UpdateByEmail(oldEmail, sameUser1)
	if err != nil {
		t.Fatal(err)
	}

	if newUser1.Email == oldEmail {
		t.Fatalf("expected different emails got: %v == %v", sameUser1.Email, newUser1.Email)
	}

	// hide password
	newUser1.Password = []byte("")

	t.Logf("%+v\n", newUser1)

	// delete test data
	err = db.Unscoped().Delete(&user1).Error
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateByID(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Fatal(err)
	}
	// migrate
	err = db.AutoMigrate(new(usermodel.User)).Error
	if err != nil {
		t.Fatal(err)
	}

	store := New(db)

	user1 := usermodel.User{
		Email:    "user1@example.com",
		Password: []byte("user1"),
	}

	sameUser1, err := store.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	// save old email
	oldEmail := sameUser1.Email
	//change email
	sameUser1.Email = "newuser1@example.com"
	sameUser1.Password = []byte("newuser1")

	newUser1, err := store.UpdateByID(sameUser1.ID, sameUser1)
	if err != nil {
		t.Fatal(err)
	}

	if newUser1.Email == oldEmail {
		t.Fatalf("expected different emails got: %v == %v", sameUser1.Email, newUser1.Email)
	}

	// hide password
	newUser1.Password = []byte("")

	t.Logf("%+v\n", newUser1)

	// delete test data
	err = db.Unscoped().Delete(&user1).Error
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {

}
