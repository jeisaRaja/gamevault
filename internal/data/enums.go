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

type Platforms string

const (
	Playstation    Platforms = "playstation"
	Xbox           Platforms = "xbox"
	NintentoSwitch Platforms = "nintento switch"
	PC             Platforms = "pc"
	IOS            Platforms = "ios"
	Android        Platforms = "android"
)

func ValidPlatform(platform Platforms) bool {
	validPlatforms := map[Platforms]bool{
		Playstation:    true,
		Xbox:           true,
		NintentoSwitch: true,
		PC:             true,
		IOS:            true,
		Android:        true,
	}
	return validPlatforms[platform]
}

func ValidateGamePlatforms(platforms []string) bool {
	for _, platform := range platforms {
		if !ValidPlatform(Platforms(platform)) {
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

func ValidateGameGenres(genres []string) bool {
	for _, genre := range genres {
		if !ValidGenre(Genres(genre)) {
			return false
		}
	}
	return true
}
