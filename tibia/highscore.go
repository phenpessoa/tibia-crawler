package tibia

import (
	"strconv"
	"strings"
)

// HighscoreCategoryFromString converts a string representation of a Highscore
// category to its corresponding HighscoreCategory.
//
// This conversion allows you to work with Highscore categories in a more
// convenient and type-safe manner.
//
// The function performs a case-insensitive comparison of the provided string
// against known Highscore categories values. If a match is found, the
// corresponding HighscoreCategory is returned along with a nil error.
//
// If the provided string does not match any known Highscore categories values,
// an ErrUnknownHighscoreCategory is returned.
//
// Strings representing the integer value of a Highscore category (i.e. "1" for
// Achievements) will also be parsed into their corresponding HighscoreCategory.
func HighscoreCategoryFromString(category string) (HighscoreCategory, error) {
	switch strings.ToLower(category) {
	case "achievement", "achievements":
		return HighscoreCategoryAchievements, nil
	case "axe", "axe fighting", "axefighting":
		return HighscoreCategoryAxeFighting, nil
	case "charm points", "charm", "charmpoints":
		return HighscoreCategoryCharmPoints, nil
	case "club fighting", "club", "clubfighting":
		return HighscoreCategoryClubFighting, nil
	case "distance fighting", "distance", "distancefighting":
		return HighscoreCategoryDistanceFighting, nil
	case "experience points", "exp", "xp", "experience", "experiencepoints",
		"level":
		return HighscoreCategoryExperiencePoints, nil
	case "fishing":
		return HighscoreCategoryFishing, nil
	case "fist fighting", "fist", "fistfighting":
		return HighscoreCategoryFistFighting, nil
	case "goshnars taint", "taint", "goshnar's taint", "goshnars",
		"goshnarstaint":
		return HighscoreCategoryGoshnarsTaint, nil
	case "loyalty points", "loyalty", "loyaltypoints":
		return HighscoreCategoryLoyaltyPoints, nil
	case "magic level", "magic", "ml", "magiclevel":
		return HighscoreCategoryMagicLevel, nil
	case "shielding":
		return HighscoreCategoryShielding, nil
	case "sword fighting", "sword", "swordfighting":
		return HighscoreCategorySwordFighting, nil
	case "drome score", "drome", "dromescore":
		return HighscoreCategoryDromeScore, nil
	case "boss points", "bosspoints":
		return HighscoreCategoryBossPoints, nil
	default:
		return HighscoreCategoryDefault, ErrUnknownHighscoreCategory
	}
}

// HighscoreCategoryFromInt converts an integer representation of a Highscore
// category to its corresponding HighscoreCategory.
//
// This conversion allows you to work with Highscore categories in a more
// convenient and type-safe manner.
//
// The function performs a comparison of the provided integer against known
// Highscore categories values. If a match is found, the corresponding
// HighscoreCategory is returned along with a nil error.
//
// If the provided integer does not match any known Highscore categories values,
// an ErrUnknownHighscoreCategory is returned.
func HighscoreCategoryFromInt(category int) (HighscoreCategory, error) {
	switch category {
	case 1:
		return HighscoreCategoryAchievements, nil
	case 2:
		return HighscoreCategoryAxeFighting, nil
	case 3:
		return HighscoreCategoryCharmPoints, nil
	case 4:
		return HighscoreCategoryClubFighting, nil
	case 5:
		return HighscoreCategoryDistanceFighting, nil
	case 6:
		return HighscoreCategoryExperiencePoints, nil
	case 7:
		return HighscoreCategoryFishing, nil
	case 8:
		return HighscoreCategoryFistFighting, nil
	case 9:
		return HighscoreCategoryGoshnarsTaint, nil
	case 10:
		return HighscoreCategoryLoyaltyPoints, nil
	case 11:
		return HighscoreCategoryMagicLevel, nil
	case 12:
		return HighscoreCategoryShielding, nil
	case 13:
		return HighscoreCategorySwordFighting, nil
	case 14:
		return HighscoreCategoryDromeScore, nil
	case 15:
		return HighscoreCategoryBossPoints, nil
	default:
		return HighscoreCategoryDefault, ErrUnknownHighscoreCategory
	}
}

// HighscoreCategory represents a Highscore category on tibia.com
type HighscoreCategory struct {
	hs int
}

var (
	// HighscoreCategoryDefault is the same as
	// HighscoreCategoryExperiencePoints.
	HighscoreCategoryDefault = HighscoreCategory{0}

	// HighscoreCategoryAchievements represents the Achievements category.
	HighscoreCategoryAchievements = HighscoreCategory{1}

	// HighscoreCategoryAxeFighting represents the Axe Fighting category.
	HighscoreCategoryAxeFighting = HighscoreCategory{2}

	// HighscoreCategoryCharmPoints represents the Charm Points category.
	HighscoreCategoryCharmPoints = HighscoreCategory{3}

	// HighscoreCategoryClubFighting represents the Club Fighting category.
	HighscoreCategoryClubFighting = HighscoreCategory{4}

	// HighscoreCategoryDistanceFighting represents the Distance Fighting
	// category.
	HighscoreCategoryDistanceFighting = HighscoreCategory{5}

	// HighscoreCategoryExperiencePoints represents the Experience Points
	// category.
	HighscoreCategoryExperiencePoints = HighscoreCategory{6}

	// HighscoreCategoryFishing represents the Fishing category.
	HighscoreCategoryFishing = HighscoreCategory{7}

	// HighscoreCategoryFistFighting represents the Fist Fighting category.
	HighscoreCategoryFistFighting = HighscoreCategory{8}

	// HighscoreCategoryGoshnarsTaint represents the Goshnar's Taint category.
	HighscoreCategoryGoshnarsTaint = HighscoreCategory{9}

	// HighscoreCategoryLoyaltyPoints represents the Loyalty Points category.
	HighscoreCategoryLoyaltyPoints = HighscoreCategory{10}

	// HighscoreCategoryMagicLevel represents the Magic Level category.
	HighscoreCategoryMagicLevel = HighscoreCategory{11}

	// HighscoreCategoryShielding represents the Shielding category.
	HighscoreCategoryShielding = HighscoreCategory{12}

	// HighscoreCategorySwordFighting represents the Sword Fighting category.
	HighscoreCategorySwordFighting = HighscoreCategory{13}

	// HighscoreCategoryDromeScore represents the Drome Score category.
	HighscoreCategoryDromeScore = HighscoreCategory{14}

	// HighscoreCategoryBossPoints represents the Boss Points category.
	HighscoreCategoryBossPoints = HighscoreCategory{15}
)

// ID returns the integer representation of the Highscore Category.
//
// It can be used to access the numerical representation of the Highscore
// Category when needed.
func (hs HighscoreCategory) ID() int {
	if hs.hs == 0 {
		return HighscoreCategoryExperiencePoints.hs
	}
	return hs.hs
}

// QueryVal returns the query parameter value representation of the Highscore
// category.
//
// The QueryVal method returns the string representation of the Highscore
// category, suitable for use as a query parameter value when making requests to
// tibia.com endpoints that support filtering by a Highscore category.
//
// Example usage:
//
//	hs := tibia.HighscoreCategoryGoshnarsTaint
//	var vals url.Values
//	vals.Set(hs.QueryKey(), hs.QueryVal())
func (hs HighscoreCategory) QueryVal() string {
	return strconv.FormatInt(int64(hs.ID()), 10)
}

// QueryKey returns the query parameter key for filtering by Highscore category.
//
// The QueryKey method returns the string representation of the query parameter
// key to be used when filtering by a Highscore category in tibia.com requests.
// This key can be appended to the query string to specify the desired
// Highscore category.
//
// Example usage:
//
//	hs := tibia.HighscoreCategoryCharmPoints
//	var vals url.Values
//	vals.Set(hs.QueryKey(), hs.QueryVal())
func (hs HighscoreCategory) QueryKey() string {
	return "category"
}

// String returns the string representation of the Highscore category.
func (hs HighscoreCategory) String() string {
	switch hs {
	case HighscoreCategoryDefault:
		return "Experience Points"
	case HighscoreCategoryAchievements:
		return "Achievements"
	case HighscoreCategoryAxeFighting:
		return "Axe Fighting"
	case HighscoreCategoryCharmPoints:
		return "Charm Points"
	case HighscoreCategoryClubFighting:
		return "Club Fighting"
	case HighscoreCategoryDistanceFighting:
		return "Distance Fighting"
	case HighscoreCategoryExperiencePoints:
		return "Experience Points"
	case HighscoreCategoryFishing:
		return "Fishing"
	case HighscoreCategoryFistFighting:
		return "Fist Fighting"
	case HighscoreCategoryGoshnarsTaint:
		return "Goshnar's taint"
	case HighscoreCategoryLoyaltyPoints:
		return "Loyalty Points"
	case HighscoreCategoryMagicLevel:
		return "Magic Level"
	case HighscoreCategoryShielding:
		return "Shielding"
	case HighscoreCategorySwordFighting:
		return "Sword Fighting"
	case HighscoreCategoryDromeScore:
		return "Drome Score"
	case HighscoreCategoryBossPoints:
		return "Boss Points"
	default:
		panic("unknown hs")
	}
}
