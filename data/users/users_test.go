package users_test

import (
	"fmt"
	"github.com/annakallo/parmtracker/data/users"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserInsertAndFetch(t *testing.T) {
	u := users.User{
		UserName:  "Orion",
		Password:  "pass",
		Email:     "orion@gmail.com",
		UserColor: "red",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := u.Insert()
	assert.Nil(t, e)
	users := users.GetUsers()
	assert.NotEqual(t, len(users), 0)
	e = u.Delete()
	assert.Nil(t, e)
}

func TestUserSave(t *testing.T) {
	u := users.User{
		UserName:  "Orion2",
		Password:  "pass1",
		Email:     "orion2@gmail.com",
		UserColor: "blue",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := u.Insert()
	assert.Nil(t, e)
	u.UserName = "Atik"
	u.Save()
	assert.Equal(t, u.UserName, "Atik")
	e = u.Delete()
	assert.Nil(t, e)
}

func TestUsersSaveSeedData(t *testing.T) {
	table := []users.User{
		{UserName: "Orion", Password: "12345", Email: "orion@gmail.com", UserColor: "#666"},
		{UserName: "Atik", Password: "11111", Email: "atik@gmail.com", UserColor: "#f0f0f0"}}

	for _, row := range table {
		row.Insert()
	}
	entries := users.GetUsers()
	assert.Equal(t, len(entries), 2)
}

func TestUsersLoadByName(t *testing.T) {
	u := users.User{}
	e := u.LoadByName("orion@gmail.com")
	if u.UserName == "" {
		e := u.LoadByEmail("orion@gmail.com")
		fmt.Println("error1", e)
	}
	fmt.Println("email", u.Email)
	fmt.Println("error2", e)
	//assert.Equal(t, len(entries), 2)
}

//func TestGetUsers(t *testing.T) {
//	users := users.GetUsers()
//	fmt.Println(users)
//	for _, u := range users {
//		fmt.Printf("%+v\n", u)
//	}
//}
