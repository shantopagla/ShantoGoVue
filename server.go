package main

import (
	"github.com/shantopagla/ShantoGoVue/api"
	"github.com/shantopagla/ShantoGoVue/models"
	"github.com/shantopagla/ShantoGoVue/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
