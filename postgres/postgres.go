package postgres

import (
	"context"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/daylightdata/shortbread/config"
)

var connection *sqlx.DB

func Connect() (*sqlx.DB, error) {

	pgURL := config.GetConfig().DatabaseURL
	conn, err := sqlx.Open("postgres", pgURL)
	if err == nil {
		conn.SetMaxOpenConns(5)
		err2 := conn.Ping()
		if err2 != nil {
			return nil, err2
		}
	}

	connection = conn

	return connection, err
}

func Connection(context.Context) *sqlx.DB {
	return connection
}

// buffalo middleware
// REMOVE THIS later
func Middleware() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			tx, err := connection.Beginx()
			if err != nil {
				return err
			}

			c.Set("db", tx)

			err = next(c)
			if err != nil {
				err2 := tx.Rollback()
				if err2 != nil {
					log.Fatal("error rolling back transaction")
				}

				return err
			}

			err = tx.Commit()
			return err
		}
	}
}
