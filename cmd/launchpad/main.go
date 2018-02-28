package main

import (
	"flag"
	"log"

	"github.com/prashantssingh/launchpad/launcher"
	"github.com/prashantssingh/launchpad/models"
)

const (
	defaultLaunchPadCount    = 2
	defaultPerPadLaunchCount = 4
	defaultSatelliteCount    = 200
)

var (
	lauchPadCount     = flag.Int("launchpadcount", defaultLaunchPadCount, "enter launchpad count")
	perPadLaunchCount = flag.Int("perpadlaunchcount", defaultPerPadLaunchCount, "enter statellite count to launch per pad in one go")
	satellitecount    = flag.Int("satellitecount", defaultSatelliteCount, "enter total count of satellites to be launched")
)

func main() {
	flag.Parse()

	lauchSettings := models.NewSatelliteSettings(*lauchPadCount, *perPadLaunchCount, *satellitecount)
	if err := launcher.Launcher(lauchSettings); err != nil {
		log.Printf("error encountered while launching: %+v", err)
	}
}
