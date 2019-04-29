package cli

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urfave/cli"
	"github.com/ysv/pick/pkg/models"
)

var sitesCmd = cli.Command{
	Name:   "sites",
	Usage:  "manage sites",
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
	},
}

func siteAdd(c *cli.Context) error {
	rand.Seed(time.Now().Unix())
	sname := c.String("site-name")
	sid := generateTrackingID()
	s := &models.Site{
		TrackingID: sid,
		Name: sname,
	}

	if err := app.database.SaveSite(s); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert the next tracking snippet into your HTML")
	printTrackingSnippet(sid)
	fmt.Println()
	return nil
}

func printTrackingSnippet(sid string) {
	fmt.Printf(`
<!-- Pick - Lightning and reliable website analytics. Let's pick it! - https://github.com/ysv/pick -->
<script>
	(function(f, a, t, h, o, m){
	a[h]=a[h]||function(){
		(a[h].q=a[h].q||[]).push(arguments)
	};
	o=f.createElement('script'),
		m=f.getElementsByTagName('script')[0];
	o.async=1; o.src=t; o.id='pick-script';
	m.parentNode.insertBefore(o,m)
})(document, window, '//localhost:8081/pick.js', 'pick');
pick('set', 'siteId', '%s');
pick('trackPageview');
</script>
<!-- / Fathom -->`, sid)
}

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
