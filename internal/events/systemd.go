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
	"github.com/arnumina/swag/util/systemd"
	"github.com/robfig/cron/v3"
)

type sdEvent struct{}

// Run AFAIRE
func (e *sdEvent) Run() {
	systemd.NotifyWatchdog()
}

// AddSystemdEvent AFAIRE
func AddSystemdEvent(c *cron.Cron) error {
	delay, err := systemd.WatchdogDelay()
	if err != nil {
		return err
	}

	if delay == 0 {
		return nil
	}

	_, err = c.AddJob("@every "+delay.String(), &sdEvent{})
	if err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
