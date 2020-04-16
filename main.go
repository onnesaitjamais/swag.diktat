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

package main

import (
	"os"

	"github.com/arnumina/swag.diktat/cmd/diktat"
)

var version, builtAt string

func main() {
	if diktat.Run(version, builtAt) != nil {
		os.Exit(-1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
