package actions

import (
	"time"

	"github.com/daylightdata/shortbread/postgres"
	"github.com/daylightdata/shortbread/spots"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_shortbread_session",
		})
		// Automatically redirect to SSL
		app.Use(forceSSL())

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		_, err := postgres.Connect()
		if err != nil {
			app.Stop(err)
		}
		app.Use(postgres.Middleware())

		// Setup and use translations:
		app.Use(translations())

		app.GET("/residents", GetAllResidents)
		app.GET("/spotstatus/{floor}", getSpots)
		app.POST("/book-spot", bookSpot)

		app.ServeFiles("/assets", assetsBox) // serve files from the public directory
		app.GET("/", HomeHandler)
		app.GET("/{path:.+}", HomeHandler)
	}

	return app
}

type spot struct {
	Space      string    `json:"space" db:"space"`
	Status     string    `json:"status" db:"status"`
	Floor      string    `json:"floor" db:"floor"`
	OccupantId uuid.UUID `json:"occupantId" db:"occupant_id"`
	Duration   int       `json:"duration" db:"duration"`
}
type currentResident struct {
	OccupantId uuid.UUID `json:"occupantId" db:"occupant_id"`
}

type spotStat struct {
	Floor  string `db:"floor_label"`
	Space  string `db:"space_label"`
	Status string `db:"status"`
}

func getSpots(c buffalo.Context) error {
	db := c.Value("db").(*sqlx.Tx)

	floorLabel := c.Param("floor")
	spotConfig := spots.All[floorLabel]

	q := `select floor_label, space_label, status from current_parking_occupancy
		where floor_label=$1`
	stats := []spotStat{}
	err := db.Select(&stats, q, floorLabel)
	if err != nil {
		return err
	}

	for _, s := range stats {
		for _, sc := range spotConfig {
			if s.Space == sc.Space {
				sc.Status = s.Status
			}
		}
	}

	return c.Render(200, r.JSON(spotConfig))
}

func bookSpot(c buffalo.Context) error {
	db := c.Value("db").(*sqlx.Tx)

	cs := spot{}
	err := c.Bind(&cs)
	if err != nil {
		return err
	}
	time.LoadLocation("Local")
	now := time.Now()
	start := now.Add(time.Hour * -4)
	x := cs.Duration
	diff := time.Minute * time.Duration(x)
	expiresAt := start.Add(diff)
	ins := `insert into current_parking_occupancy (occupant_id, floor_label, space_label, status, started_at, expires_at)
			values ($1,$2,$3,$4,$5,$6)`

	_, err = db.Exec(ins, cs.OccupantId, cs.Floor, cs.Space, "occupied", start, expiresAt)
	if err != nil {
		return err
	}

	return c.Render(200, nil)
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
