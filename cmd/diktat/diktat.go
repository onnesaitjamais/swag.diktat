/*
#######
##                                     ___ __   __       __
##        ____    _____ ____ _     ___/ (_) /__/ /____ _/ /_
##       (_-< |/|/ / _ `/ _ `/ _  / _  / /  '_/ __/ _ `/ __/
##      /___/__,__/\_,_/\_, / (_) \_,_/_/_/\_\\__/\_,_/\__/
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package diktat

import (
	"github.com/arnumina/swag"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/options"
	"github.com/robfig/cron/v3"

	"github.com/arnumina/swag.diktat/internal/events"
)

func initialize(s *service.Service) error {
	cron := cron.New(cron.WithSeconds())

	if err := events.AddSystemdEvent(cron); err != nil {
		return err
	}

	atStartup, err := events.AddCronEvents(s, cron)
	if err != nil {
		return err
	}

	s.AddGroupFn(
		func(stop <-chan struct{}) error {
			for _, e := range atStartup {
				e.Run()
			}

			atStartup = nil

			cron.Start()
			<-stop
			cron.Stop()

			return nil
		},
	)

	return nil
}

// Run AFAIRE
func Run(version, builtAt string) error {
	service, err := swag.NewService(
		"diktat",
		version,
		builtAt,
		swag.Broker(
			"mongodb",
			options.Options{},
		),
	)
	if err != nil {
		return err
	}

	defer service.Close()

	if err := initialize(service); err != nil {
		service.Logger().Critical( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Error when initializing this service",
			"name", service.Name(),
			"version", version,
			"reason", err.Error(),
		)

		return err
	}

	if err := service.Run(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
