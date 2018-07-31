package grifts

import (
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/markbates/grift/grift"

	"github.com/daylightdata/shortbread/postgres"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		db, err := postgres.Connect()
		if err != nil {
			return err
		}
		defer db.Close()

		err = fake.SetLang("en")
		if err != nil {
			return err
		}

		ins := `insert into resident_information (id, first_name, last_name, phone_number, room_number)
			values ($1,$2,$3,$4,$5)`

		for i := 0; i < 100; i++ {
			_, err := db.Exec(ins, uuid.New(), fake.FirstName(), fake.LastName(), fake.Phone(), fake.Zip())
			if err != nil {
				return err
			}
		}

		return nil
	})

})
