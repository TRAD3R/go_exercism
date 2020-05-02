package space

type Planet string

func Age(seconds float64, planet Planet) (age float64) {
	switch planet {
		case "Mercury":
			age = seconds / 31557600 / 0.2408467
		case "Venus":
			age = seconds / 31557600 / 0.61519726
		case "Earth":
			age = seconds / 31557600
		case "Mars":
			age = seconds / 31557600 / 1.8808158
		case "Jupiter":
			age = seconds / 31557600 / 11.862615
		case "Saturn":
			age = seconds / 31557600 / 29.447498
		case "Uranus":
			age = seconds / 31557600 / 84.016846
		case "Neptune":
			age = seconds / 31557600 / 164.79132
	}

	return age
}
