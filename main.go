package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Planet struct {
	Name             string
	Diameter         float64
	Gravity          float64
	Mass             float64
	Type             string
	Livable          bool
	Description      string
	Atmosphere       string
	Moons            int
	Temperature      float64
	RotationPeriod   float64
	RevolutionPeriod float64
	SurfaceFeatures  []string
	MagneticField    bool
	Rings            bool
	Exploration      string
}

var planetNames = []string{
	"Novus",
	"Stella",
	"Astraea",
	"Epsilon",
	"Orionis",
	"Nebula",
	"Seraphina",
	"Zenith",
	"Arctica",
	"Valora",
	"Solstice",
	"Galactron",
	"Nebulon",
	"Celestis",
	"Zirconia",
	"Helios",
	"Aquila",
	"Eridanus",
	"Cassiopeia",
	"Lunaris",
	"Aetherius",
	"Calypso",
	"Astralis",
	"Volantis",
	"Lyricus",
	"Phobos",
	"Isolde",
	"Solara",
	"Electra",
	"Arcanum",
	"Orionis",
	"Zephyrus",
	"Andromeda",
	"Arcturus",
	"Astraia",
	"Nocturna",
	"Vespera",
	"Nebulus",
	"Solis",
	"Helix",
	"Galaxia",
	"Vega",
	"Nova",
	"Draco",
	"Bellatrix",
	"Lumina",
	"Celestia",
	"Elysia",
	"Aurius",
	"Selene",
}
var planetDescriptions = []string{
	"This planet is known for its breathtaking landscapes and vibrant ecosystems.",
	"As one of the largest planets in the galaxy, it boasts majestic rings and numerous moons.",
	"This planet is a frozen world with extreme temperatures and icy terrains.",
	"Home to stunning auroras and powerful magnetic fields, this planet is a sight to behold.",
	"Scientists believe this planet may have the potential to support life due to its favorable conditions.",
	"With its vast oceans and lush forests, this planet is a haven for biodiversity.",
	"Featuring towering mountains and vast deserts, this planet is a study in contrasting landscapes.",
	"This planet is shrouded in mystery, with ancient ruins and unexplored territories waiting to be discovered.",
	"Known for its extreme weather phenomena, this planet experiences frequent storms and atmospheric turbulence.",
	"Rich in valuable resources, this planet has attracted the attention of intergalactic miners and explorers.",
}
var planetAtmospheres = []string{
	"Primarily composed of nitrogen and oxygen, this planet's atmosphere supports various life forms.",
	"This planet has a dense atmosphere rich in hydrogen and helium, making it an ideal gas giant.",
	"The atmosphere of this planet is composed of methane and ammonia, resulting in its distinctive blue appearance.",
	"With a thin atmosphere of carbon dioxide, this planet experiences extreme temperature variations.",
	"Scientists have yet to fully explore and understand the complex atmosphere of this exoplanet.",
	"The atmosphere of this planet is laden with sulfur compounds, creating a pungent and acidic environment.",
	"Featuring swirling clouds of dust and gas, this planet's atmosphere gives it an ethereal beauty.",
	"The atmosphere of this planet is known for its mesmerizing light displays and atmospheric electricity.",
	"This planet experiences frequent atmospheric disturbances, resulting in intense auroras and atmospheric phenomena.",
	"The atmosphere of this planet contains rare elements that give it a unique coloration and composition.",
}

func main() {
	// Generate a seed based on current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random planet
	planet := generateRandomPlanet()

	// Print the planet information
	fmt.Println("")
	fmt.Println("Planet Name:", planet.Name)
	fmt.Println("Planet Type:", planet.Type)
	fmt.Println("Description:", planet.Description)
	fmt.Println("Surface Features:", planet.SurfaceFeatures)
	fmt.Println("Livable:", planet.Livable)
	fmt.Println("Temperature:", planet.Temperature, "°C")

	fmt.Printf("\r\n######### INFO #########\r\n")

	fmt.Println("Diameter:", planet.Diameter, "km")
	fmt.Println("Gravity:", planet.Gravity, "m/s^2")
	fmt.Println("Mass:", planet.Mass, "kg")
	fmt.Println("Atmosphere:", planet.Atmosphere)
	fmt.Println("Moons:", planet.Moons)
	fmt.Println("Rotation Period:", planet.RotationPeriod, "hours")
	fmt.Println("Revolution Period:", planet.RevolutionPeriod, "days")
	fmt.Println("Magnetic Field:", planet.MagneticField)
	fmt.Println("Rings:", planet.Rings)
	fmt.Println("Exploration Status:", planet.Exploration)
}

func generateRandomPlanet() Planet {
	// Generate a random index for the planetTypes, planetNames, planetDescriptions, and planetAtmospheres slices
	nameIndex := rand.Intn(len(planetNames))
	descriptionIndex := rand.Intn(len(planetDescriptions))
	atmosphereIndex := rand.Intn(len(planetAtmospheres))

	// Generate a random mass between 1e22 and 1e28 kg
	mass := rand.Float64()*(1e28-1e22) + 1e22

	// Calculate the diameter based on the mass
	diameter := calculateDiameter(mass)

	// Calculate the gravity based on the mass and diameter
	gravity := calculateGravity(mass, diameter)

	// Generate random values for other planet properties
	moons := rand.Intn(10)
	temperature := rand.Float64()*200 - 100  // Random temperature between -100°C and 100°C
	rotationPeriod := rand.Float64() * 24    // Random rotation period between 0 and 24 hours
	revolutionPeriod := rand.Float64() * 100 // Random revolution period between 0 and 100 days

	// Generate a random magnetic field status (true or false)
	magneticField := rand.Intn(2) == 1

	// Determine livable status based on temperature
	livable := isPlanetLivable(magneticField, temperature)

	// Generate random surface features based on livable status
	surfaceFeatures := generateRandomSurfaceFeatures(livable)

	// Generate a random rings status (true or false)
	rings := rand.Intn(2) == 1

	// Generate a random exploration status
	explorationStatus := generateRandomExplorationStatus()

	planetType := determinePlanetType(livable, temperature)

	// Create and return the Planet struct
	planet := Planet{
		Name:             planetNames[nameIndex],
		Diameter:         diameter,
		Gravity:          gravity,
		Mass:             mass,
		Type:             planetType,
		Livable:          livable,
		Description:      planetDescriptions[descriptionIndex],
		Atmosphere:       planetAtmospheres[atmosphereIndex],
		Moons:            moons,
		Temperature:      temperature,
		RotationPeriod:   rotationPeriod,
		RevolutionPeriod: revolutionPeriod,
		SurfaceFeatures:  surfaceFeatures,
		MagneticField:    magneticField,
		Rings:            rings,
		Exploration:      explorationStatus,
	}

	return planet
}

func calculateDiameter(mass float64) float64 {
	// The calculation formula can be adjusted based on scientific models or assumptions
	return mass / 5e20
}

func calculateGravity(mass, diameter float64) float64 {
	// The calculation formula can be adjusted based on scientific models or assumptions
	return (6.67430e-11 * mass) / (0.5 * diameter * diameter)
}

func isPlanetLivable(magneticField bool, temperature float64) bool {

	if !magneticField {
		return false
	}

	// Define the livable temperature range
	minTemperature := -20.0
	maxTemperature := 50.0

	// Check if the temperature falls within the livable range
	return temperature >= minTemperature && temperature <= maxTemperature
}

func generateRandomSurfaceFeatures(livable bool) []string {
	var surfaceFeatures []string

	if livable {
		surfaceFeatures = []string{
			"Mountains",
			"Valleys",
			"Oceans",
			"Forests",
			"Grassland",
			"Swamps",
			"Lakes",
		}
	} else {
		surfaceFeatures = []string{
			"Craters",
			"Deserts",
			"Volcanoes",
			"Canyons",
			"Ice Caps",
			"Hurricanes",
			"Tornados",
			"Caves",
		}
	}

	rand.Shuffle(len(surfaceFeatures), func(i, j int) {
		surfaceFeatures[i], surfaceFeatures[j] = surfaceFeatures[j], surfaceFeatures[i]
	})

	numFeatures := rand.Intn(5) + 1 // Random number of features between 1 and 5
	return surfaceFeatures[:numFeatures]
}

func generateRandomExplorationStatus() string {
	explorationStatus := []string{
		"Unexplored",
		"Partially Explored",
		"Extensively Explored",
		"Inhabited",
		"Colonized",
	}

	index := rand.Intn(len(explorationStatus))
	return explorationStatus[index]
}

func determinePlanetType(livable bool, temperature float64) string {
	if livable {
		if temperature < -30 {
			return "Ice Planet"
		} else if temperature > 30 {
			return "Desert Planet"
		} else {
			return "Terrestrial Planet"
		}
	} else {
		return "Gas Giant"
	}
}
