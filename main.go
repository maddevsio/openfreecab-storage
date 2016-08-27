package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-storage/conf"
	"github.com/maddevsio/openfreecab-storage/service"
	"github.com/urfave/cli"
)

func main() {
	app := conf.NewConfigurator()
	app.App().Action = func(c *cli.Context) error {
		worker := service.NewOpenStorage(app.Get())
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, os.Kill)
		defer signal.Stop(signalChan)
		go func() {
			<-signalChan
			log.Info("Interrupting...")
			worker.Stop()
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}()
		err := worker.Start()
		if err != nil {
			log.Fatalf("error starting service, %v", err)
		}
		worker.WaitStop()
		return nil
	}
	if err := app.Run(); err != nil {
		log.Fatalf("error on run app, %v", err)
	}
}
