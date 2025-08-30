package space

type Planet string

const earthYear = 31557600 // in seconds

func Age(seconds float64, planet Planet) float64 {
	pYear := map[Planet]float64{"Mercury":0.2408467,
                            "Venus":0.61519726,
                            "Earth":1.0,
                            "Mars":1.8808158,
                            "Jupiter":11.862615,
                            "Saturn":29.447498,
                            "Uranus":84.016846,
                            "Neptune":164.79132}
    earthAge := seconds / earthYear
    if _, ok := pYear[planet]; !ok {
        return -1.0
    }
    return earthAge / pYear[planet]
}
