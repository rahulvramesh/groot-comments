package models

import (
	"fmt"
	"math/rand"

	"github.com/rahulvramesh/groot-comments/db"

	"github.com/bxcodec/faker"
)

func Seed() {

	db := db.GetSession()

	//for members
	n := 0
	for n < 11 {

		a := Member{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}

		a.ID = uint(n)
		fmt.Println(a)
		db.Create(&a)
		n++
	}

	//for comments
	n = 0
	for n < 5 {

		b := Comment{}
		err := faker.FakeData(&b)
		if err != nil {
			fmt.Println(err)
		}

		b.ID = uint(n)
		b.Organization = "xendit"
		// rando := rand.Seed(10)
		b.AuthorID = rand.Intn(10)
		fmt.Println(b)
		db.Create(&b)
		n++
	}
}
