package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	null "gopkg.in/guregu/null.v3"
)

type resident struct {
	Id        uuid.UUID   `json:"id" db:"id"`
	FirstName string      `json:"firstName" db:"first_name"`
	LastName  string      `json:"lastName" db:"last_name"`
	Phone     string      `json:"phone" db:"phone_number"`
	Room      string      `json:"room" db:"room_number"`
	Spot      null.String `json:"spot" db:"space_label"`
	Floor     null.String `json:"floor" db:"floor_label"`
}

func GetAllResidents(c buffalo.Context) error {
	db := c.Value("db").(*sqlx.Tx)
	allResidents := []resident{}

	q := `select
		ri.id, ri.first_name, ri.last_name, ri.phone_number, ri.room_number, cpo.floor_label, cpo.space_label
		from resident_information ri left join current_parking_occupancy cpo
			on ri.id=cpo.occupant_id`
	err := db.Select(&allResidents, q)
	if err != nil {
		return errors.Wrap(err, "selecting all residents")
	}

	return c.Render(200, r.JSON(allResidents))
}
