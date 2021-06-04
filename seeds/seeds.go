package seeds

import "go.uber.org/fx"

// Module exports seed module
var Module = fx.Options(
	fx.Provide(NewSuperAdminUser),
	fx.Provide(NewSeeds),
)

// Seed db seed
type Seed interface {
	Run()
}

// Seeds listing of seeds
type Seeds []Seed

// Run -> runs the seed data
func (s Seeds) Run() {
	for _, seed := range s {
		seed.Run()
	}
}

// NewSeeds -> creates new seeds
func NewSeeds(
	superAdminSeed SuperAdminSeed,
) Seeds {
	return Seeds{
		superAdminSeed,
	}
}
