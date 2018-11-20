package space

const earthYearInSeconds = 31557600.0

type Planet string

var orbitTimeInEarthYears = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	//try to find planet in planet dictionary
	if orbitTime, found := orbitTimeInEarthYears[planet]; !found {
		return (-1.0)
	} else {
		secondsInPlanetYear := orbitTime * earthYearInSeconds
		return (seconds / secondsInPlanetYear)
	}
}
