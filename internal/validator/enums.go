package validator

func EnumToStringSlice[T ~string](le []T) []string {
	result := make([]string, len(le))
	for i, v := range le {
		result[i] = string(v)
	}
	return result
}

type Genres string

const (
	Action          Genres = "action"
	Adventure       Genres = "adventure"
	ActionAdventure Genres = "action-adventure"
	Puzzle          Genres = "puzzle"
	RolePlaying     Genres = "role-playing"
	Simulation      Genres = "simulation"
	Strategy        Genres = "strategy"
	Sports          Genres = "sports"
	MMO             Genres = "mmo"
	Platformer      Genres = "platformer"
)

type Platform string

const (
	Playstation4   Platform = "playstation4"
	Playstation5   Platform = "playstation5"
	XboxOne        Platform = "xbox one"
	XboxS          Platform = "xbox series s"
	NintentoSwitch Platform = "nintento switch"
	Steam          Platform = "steam"
	Epic           Platform = "epic"
	IOS            Platform = "ios"
	Android        Platform = "android"
)
