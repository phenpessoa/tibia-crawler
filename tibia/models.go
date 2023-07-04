package tibia

// BoostableBoss represents information about a boostable boss.
//
// The BoostableBoss struct contains details about a specific boss, including
// its name, the URL to the boss's image, and whether the boss is boosted today
// or not. This information is typically obtained from the tibia.com Boostable
// Bosses Library page.
type BoostableBoss struct {
	// Name is the name of the boss.
	Name string `json:"name"`

	// ImageURL is the URL to the image of the boss.
	ImageURL string `json:"image_url"`

	// IsBoosted reports whether the boss is boosted today or not.
	IsBoosted bool `json:"featured"`
}

// BoostableBosses represents information about boostable bosses.
//
// The BoostableBosses struct contains information about today's boosted boss
// and a list of all the boostable bosses. It is typically used to represent the
// parsed and structured data obtained from the tibia.com Boostable Bosses
// Library page.
type BoostableBosses struct {
	// Boosted is today's boosted boss.
	Boosted BoostableBoss `json:"boosted"`

	// Bosses is a list of all boostable bosses.
	Bosses []BoostableBoss `json:"boostable_boss_list"`
}
