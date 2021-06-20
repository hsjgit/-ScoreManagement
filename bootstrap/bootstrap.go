package bootstrap

import (
	"log"
	"net/http"

	"github.com/ScoreManagement/lib/db"
	"github.com/ScoreManagement/route"
)

func Bootstrap(path, addr string) error {
	if err := db.ConnectDB(path); err != nil {
		return err
	}
	route.RegisterRoutes()
	log.Println("listen " + addr)
	return http.ListenAndServe(addr, nil)

}
