package gorm

import "testing"

type User struct {
	Name string
	Id   int64
}

func getDB() DB {
	db, _ := Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	return db
}

func TestSaveAndFirst(t *testing.T) {
	// create table "users" ("name" varchar(255));
	db := getDB()
	u := &User{Name: "jinzhu"}
	db.Save(u)
	if u.Id == 0 {
		t.Errorf("Should have ID after create record")
	}

	user := &User{}
	db.First(user)
	if user.Name != "jinzhu" {
		t.Errorf("User should be saved and fetched correctly")
	}

	users := []User{}
	db.Find(&users)
	for _, user := range users {
		if user.Name != "jinzhu" {
			t.Errorf("User should be saved and fetched correctly")
		}
	}
}

func TestWhere(t *testing.T) {
	db := getDB()
	u := &User{Name: "jinzhu"}
	db.Save(u)

	user := &User{}
	db.Where("Name = ?", "jinzhu").First(user)
	if user.Name != "jinzhu" {
		t.Errorf("Should found out user with name 'jinzhu'")
	}

	user = &User{}
	orm := db.Where("Name = ?", "jinzhu-noexisting").First(user)
	if orm.Error == nil {
		t.Errorf("Should return error when looking for unexist record, %+v", user)
	}

	users := &[]User{}
	orm = db.Where("Name = ?", "jinzhu-noexisting").First(users)
	if orm.Error != nil {
		t.Errorf("Shouldn't return error when looking for unexist records, %+v", users)
	}
}

func TestCreateTable(t *testing.T) {
	db := getDB()
	db.Exec("drop table users;")

	orm := db.CreateTable(&User{})
	if orm.Error != nil {
		t.Errorf("No error should raise when create table, but got %+v", orm.Error)
	}
}