package server

import (
	"github.com/michaelgbenle/Boiler-plate/database"
	"github.com/michaelgbenle/Boiler-plate/handlers"
	"github.com/michaelgbenle/Boiler-plate/router"
	"log"
)

func Start() error {
	values := database.InitializeDbParameters()
	var PDB = new(database.PostgresDb)
	h := &handlers.Handler{DB: PDB}

	err := PDB.SetupDb(values.Host, values.User, values.Password, values.DbName, values.Port)
	if err != nil {
		log.Fatal(err)
	}
	routes, port := router.SetupRouter(h)
	err = routes.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
