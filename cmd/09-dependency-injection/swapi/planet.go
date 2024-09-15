package swapi

type Climate string
type Gravity float64
type Biom string

type Planet struct {
	name           string
	rotationPeriod int
	orbitalPeriod  int
	diameter       int
	climate        Climate
	gravity        Gravity
	terrains       []Biom
	surfaceWater   int
	population     int
}

func NewPlanet(
	name string,
	rotationPeriod,
	orbitalPeriod,
	diameter int,
	climate Climate,
	gravity Gravity,
	terrains []Biom,
	surfaceWater,
	population int,
) *Planet {
	return &Planet{
		name:           name,
		rotationPeriod: rotationPeriod,
		orbitalPeriod:  orbitalPeriod,
		diameter:       diameter,
		climate:        climate,
		gravity:        gravity,
		terrains:       terrains,
		surfaceWater:   surfaceWater,
		population:     population,
	}
}

func (p *Planet) Name() string {
	return p.name
}

func (p *Planet) RotationPeriod() int {
	return p.rotationPeriod
}

func (p *Planet) OrbitalPeriod() int {
	return p.orbitalPeriod
}

func (p *Planet) Diameter() int {
	return p.diameter
}

func (p *Planet) Climate() Climate {
	return p.climate
}

func (p *Planet) Gravity() Gravity {
	return p.gravity
}

func (p *Planet) Terrains() []Biom {
	return p.terrains
}

func (p *Planet) SurfaceWater() int {
	return p.surfaceWater
}

func (p *Planet) Population() int {
	return p.population
}
