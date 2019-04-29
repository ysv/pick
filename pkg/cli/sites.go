package cli

import (
	"fmt"
	"math/rand"

	"github.com/urfave/cli"
	"github.com/ysv/pick/pkg/models"
)

var sitesCmd = cli.Command{
	Name:   "sites",
	Usage:  "manage sites",
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "site-id",
			Usage: "ID of the site to show for editing",
		},
		cli.StringFlag{
			Name:  "site-name",
			Usage: "",
		},
		cli.StringFlag{
			Name:  "end-date",
			Usage: "end date, expects a date in format 2006-01-02",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "get a json response",
		},
	},
	Subcommands: []cli.Command{
		{
			Name: "add",
			Action: siteAdd,
			Flags:  []cli.Flag {
				cli.StringFlag{
					Name:  "site-name",
				},
			},

		},
		{
			Name: "edit",
			Action: siteEdit,
			Flags:  []cli.Flag {
				cli.StringFlag{
					Name:  "site-name",
				},
				cli.StringFlag{
					Name:  "site-id",
				},
			},

		},

	},
}

func siteAdd(c *cli.Context) error {
	sname := c.String("site-name")
	sid := generateTrackingID()
	s := &models.Site{
		TrackingID: sid,
		Name: sname,
	}

	if err := app.database.SaveSite(s); err != nil {
		return err
	}
	fmt.Println(sid)
	return nil
}

func siteEdit(c *cli.Context) error {
	sname := c.String("site-name")
	sid := generateTrackingID()
	s := &models.Site{
		TrackingID: sid,
		Name: sname,
	}

	if err := app.database.SaveSite(s); err != nil {
		return err
	}
	fmt.Println(sid)
	return nil
}

//func saveSite(w http.ResponseWriter, r *http.Request) error {
//	var s *models.Site
//	vars := mux.Vars(r)
//	sid, ok := vars["id"]
//	if ok {
//		id, err := strconv.ParseInt(sid, 10, 64)
//		if err != nil {
//			return err
//		}
//
//		s, err = app.database.GetSite(id)
//		if err != nil {
//			return err
//		}
//	} else {
//		s = &models.Site{
//			TrackingID: generateTrackingID(),
//		}
//	}
//
//	err := json.NewDecoder(r.Body).Decode(s)
//	if err != nil {
//		return err
//	}
//
//	if err := app.database.SaveSite(s); err != nil {
//		return err
//	}
//
//	return respond(w, http.StatusOK, envelope{Data: s})
//}

func generateTrackingID() string {
	return randomString(5)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //a=65 and z = 65+25
	}

	return string(bytes)
}
