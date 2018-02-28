package launcher

import (
	"log"

	"fmt"
	"strconv"
	"sync"

	"github.com/prashantssingh/launchpad/models"
)

type satelliteSlice struct {
	satellites []string
}

var (
	doneCh = make(chan bool)
)

// Launcher takes satellite launch settings and launches them in groups,
// maxed to the capacity that can be launched per launchpad.
func Launcher(ls models.LaunchSettings) error {

	// we have to launch as many goroutines as count of launch pads
	// this is main goroutine which will wait for all launch pads to finish
	var s sync.WaitGroup
	s.Add(ls.LaunchpadCount)

	// each launch pad will launch its own satellites and signal next launch pad to start
	// For example: Launchpad1((LP1) will start first, lauch required satellites and then pass
	// the control to LP2. After LP2 finishes the lauch and pass the control to LP1.
	passControlChans := make([]chan int, ls.LaunchpadCount)

	for i := 0; i < ls.LaunchpadCount; i++ {
		passControlChans[i] = make(chan int)
	}

	for i := 0; i < ls.LaunchpadCount-1; i++ {
		go launchSatellites(i+1, passControlChans[i], passControlChans[i+1], ls.PerPadLaunchCount, ls.SatelliteCount, &s)
	}

	go launchSatellites(ls.LaunchpadCount, passControlChans[ls.LaunchpadCount-1], passControlChans[0], ls.PerPadLaunchCount, ls.SatelliteCount, &s)

	// send the satellite ID (lets say) to the first launchpad and wait for entire launch to get concluded
	passControlChans[0] <- 1

	s.Wait()

	fmt.Println("Boyyy SpaceX is done launching satellites")

	return nil
}

// Inputs:
// id: ID of launchpad such as LP1, LP2
// rcvChan: channel to receive start launchpad ID of the satellites to be launched next from the previous launchpad
// sendChan: channel to send start lauchpad ID of the satellites to be launched next to the next launch pad
func launchSatellites(id int, rcvChan chan int, sendChan chan int, launchCount int, totalCount int, s *sync.WaitGroup) {

	site := "LP" + strconv.Itoa(id)
	for {
		select {
		case start := <-rcvChan:

			// slice of satellites to be launched simulteneously
			var satNum []int
			for j := 0; j < launchCount; j++ {
				if start+j <= totalCount {
					satNum = append(satNum, start+j)
				} else {
					break
				}
			}

			preparedSatellites := prepareSatelliteForLaunch(satNum)
			log.Printf("Launching statellites from site %s... \t%+v", site, preparedSatellites.satellites)

			// if we are done launching total number satellites, tell this to all launchpads by simply closing the donechan
			start = start + launchCount
			if start >= totalCount {
				close(doneCh)
				s.Done()
				return
			}

			// continue if the
			sendChan <- start

		case <-doneCh:
			// the launchpad which finishes launching last satellites closes doneCh
			s.Done()
			return
		}
	}
}
