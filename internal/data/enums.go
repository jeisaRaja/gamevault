package data

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
	Playstation    Platform = "playstation"
	Xbox           Platform = "xbox"
	NintentoSwitch Platform = "nintento switch"
	PC             Platform = "pc"
	IOS            Platform = "ios"
	Android        Platform = "android"
)

func ValidPlatform(platform Platform) bool {
	validPlatforms := map[Platform]bool{
		Playstation:    true,
		Xbox:           true,
		NintentoSwitch: true,
		PC:             true,
		IOS:            true,
		Android:        true,
	}
	return validPlatforms[platform]
}

func ValidateGamePlatforms(platforms []Platform) bool {
	for _, v := range platforms {
		if !ValidPlatform(v) {
			return false
		}
	}

	return true
}

func ValidGenre(genre Genres) bool {
	validGenres := map[Genres]bool{
		Action:          true,
		Adventure:       true,
		ActionAdventure: true,
		Puzzle:          true,
		RolePlaying:     true,
		Simulation:      true,
		Strategy:        true,
		Sports:          true,
		MMO:             true,
		Platformer:      true,
	}

	return validGenres[genre]
}

func ValidateGameGenres(genres []Genres) bool {
	for _, genre := range genres {
		if !ValidGenre(genre) {
			return false
		}
	}
	return true
}
