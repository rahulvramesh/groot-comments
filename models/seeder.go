package models

import (
	"fmt"

	"github.com/rahulvramesh/groot-comments/db"

	"github.com/bxcodec/faker"
)

func Seed() {

	db := db.GetSession()

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

}
