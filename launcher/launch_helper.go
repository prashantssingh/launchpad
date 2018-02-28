package launcher

import (
	"fmt"
)

func prepareSatelliteForLaunch(satArr []int) satelliteSlice {
	var satStrArr []string
	for i := 0; i < len(satArr); i++ {
		satStrArr = append(satStrArr, fmt.Sprintf("%d", satArr[i]))
	}

	return satelliteSlice{
		satellites: satStrArr,
	}
}