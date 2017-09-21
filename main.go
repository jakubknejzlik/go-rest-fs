package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "s,storage",
			EnvVar: "STORAGE_URL",
		},
		cli.StringFlag{
			Name:   "d,database",
			EnvVar: "DATABASE_URL",
		},
	}
	app.Action = func(c *cli.Context) error {

		dbURL := c.String("database")
		if dbURL == "" {
			return cli.NewExitError("database flag missing", 1)
		}
		storageURL := c.String("storage")
		if storageURL == "" {
			return cli.NewExitError("storage flag missing", 1)
		}

		db, err := NewDB(dbURL)
		if err != nil {
			panic(err)
		}

		r := getRouter(db)
		log.Fatal(http.ListenAndServe(":8080", r))
	}

	app.Run(os.Args)
}
