package launcher

import (
	"log"

	"github.com/launchpad/models"
)

type satelliteSlice struct {
	satellites []string
}

var (
	satCh  = make(chan satelliteSlice)
	doneCh = make(chan bool)
	done   = true
)

// Launcher takes satellite launch settings and launches them in groups,
// maxed to the capacity that can be launched per launchpad.
func Launcher(ls models.LaunchSettings) error {

	// setup lauchpads and defer breakdown once all saatellites
	// are launched
	spawnFuncs(ls.LaunchpadCount, launchSatellites)
	defer spawnFuncs(ls.LaunchpadCount, sendDone)

	for i := 1; i <= ls.SatelliteCount; i += ls.PerPadLaunchCount {
		var satNum []int
		for j := i; j <= i+(ls.PerPadLaunchCount-1); j++ {
			satNum = append(satNum, j)
		}
		preparedSatellites := prepareSatelliteForLaunch(satNum)
		satCh <- preparedSatellites
	}

	return nil
}

func launchSatellites() {
	for {
		select {
		case sats := <-satCh:
			log.Printf("Launching statellites... \t%+v", sats.satellites)
		case <-doneCh:
			log.Println("DONE!!")
		}
	}
}

func sendDone() {
	doneCh <- done
}
