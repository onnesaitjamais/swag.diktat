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

package events

import (
	"github.com/arnumina/swag/component"
	"github.com/arnumina/swag/service"
	"github.com/robfig/cron/v3"
)

type tools struct {
	broker component.Broker
	logger component.Logger
}

type cronEvent struct {
	tools *tools
	name  string
}

// Run AFAIRE
func (e *cronEvent) Run() {
	e.tools.logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Emit",
		"event", e.name,
	)

	e.tools.broker.Publish(e.name, nil)
}

// AddCronEvents AFAIRE
func AddCronEvents(s *service.Service, c *cron.Cron) ([]cron.Job, error) {
	tools := &tools{
		broker: s.Broker(),
		logger: s.Logger(),
	}
	atStartup := []cron.Job{}

	cfg, err := s.ServiceCfg("diktat")
	if err != nil {
		return nil, err
	}

	events, err := cfg.Slice("events")
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		if disabled, err := event.DBool(false, "disabled"); err != nil {
			return nil, err
		} else if disabled {
			continue
		}

		name, err := event.String("name")
		if err != nil {
			return nil, err
		}

		e := &cronEvent{
			tools: tools,
			name:  name,
		}

		if as, err := event.DBool(false, "at_startup"); err != nil {
			return nil, err
		} else if as {
			atStartup = append(atStartup, e)
		}

		cronSpec, err := event.String("cron")
		if err != nil {
			return nil, err
		}

		_, err = c.AddJob(cronSpec, e)
		if err != nil {
			return nil, err
		}
	}

	return atStartup, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
