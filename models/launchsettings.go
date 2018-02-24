package models

// LaunchSettings stores configuration for the current
// launchpad app
type LaunchSettings struct {
	LaunchpadCount    int
	PerPadLaunchCount int
	SatelliteCount    int
}

// NewSatelliteSettings initialises the luanchpad with
// provided settings
func NewSatelliteSettings(launchpadcount, perpadlaunchcount, satellitecount int) LaunchSettings {
	return LaunchSettings{
		LaunchpadCount:    launchpadcount,
		PerPadLaunchCount: perpadlaunchcount,
		SatelliteCount:    satellitecount,
	}
}
