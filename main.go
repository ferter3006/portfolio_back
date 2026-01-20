package main

import (
	"log"
	"net/http"
	"new-test/controllers"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		// test hi world endpoint
		se.Router.GET("/hi", func(e *core.RequestEvent) error {
			return e.JSON(http.StatusOK, map[string]bool{"success": true})
		})

		se.Router.GET("/hi2", controllers.HiWorld)

		// Redsys config
		redsysCfg, err := controllers.LoadRedsysConfigFromDefault()
		if err != nil {
			return err
		}

		// Endpoint para iniciar pago
		se.Router.GET("/redsys/pay", func(e *core.RequestEvent) error {
			controllers.RedsysPayHandler(redsysCfg)(e.Response, e.Request)
			return nil
		})

		// Endpoint para notificación
		se.Router.POST("/redsys/notify", func(e *core.RequestEvent) error {
			controllers.RedsysNotifyHandler(redsysCfg)(e.Response, e.Request)
			return nil
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
