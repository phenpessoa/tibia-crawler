package tibia

import (
	"strconv"
	"strings"
)

// WorldTypeFromString converts a string representation of a World Type to its
// corresponding WorldType.
//
// This conversion allows you to work with World types in a more convenient and
// type-safe manner.
//
// The function performs a case-insensitive comparison of the provided string
// against known world types. If a match is found, the corresponding WorldType
// is returned along with a nil error.
//
// If the provided string does not match any known world types,
// an ErrUnknownWorldType is returned.
//
// Strings representing the integer value of a WorldType (i.e. "1" for
// Opptional PvP) will also be parsed into their corresponding WorldType.
func WorldTypeFromString(be string) (WorldType, error) {
	switch strings.ToLower(be) {
	case "open", "openpvp", "open pvp", "0":
		return WorldTypeOpenPvP, nil
	case "optional", "optionalpvp", "optional pvp", "1":
		return WorldTypeOptionalPvP, nil
	case "hardcore", "hardcorepvp", "hardcore pvp", "2":
		return WorldTypeHardcorePvP, nil
	case "retro open", "retroopen", "retro open pvp", "retroopenpvp", "3":
		return WorldTypeRetroOpenPvP, nil
	case "retro hardcore", "retrohardcore", "retro hardcore pvp",
		"retrohardcorepvp", "4":
		return WorldTypeRetroHardcorePvP, nil
	default:
		return WorldType{}, ErrUnknownWorldType
	}
}

// WorldTypeFromInt converts an integer representation of a World Type to its
// corresponding World Type.
//
// This conversion allows you to work with World Type in a more convenient and
// type-safe manner.
//
// The function performs a comparison of the provided integer against known
// World types. If a match is found, the corresponding World Type is returned
// along with a nil error.
//
// If the provided integer does not match any known World types,
// an ErrUnknownWorldType is returned.
func WorldTypeFromInt(be int) (WorldType, error) {
	switch be {
	case 0:
		return WorldTypeOpenPvP, nil
	case 1:
		return WorldTypeOptionalPvP, nil
	case 2:
		return WorldTypeHardcorePvP, nil
	case 3:
		return WorldTypeRetroOpenPvP, nil
	case 4:
		return WorldTypeRetroHardcorePvP, nil
	default:
		return WorldType{}, ErrUnknownWorldType
	}
}

// WorldType represents a tibia world type, such as Open PvP.
type WorldType struct {
	wt int
}

var (
	// WorldTypeOpenPvP represents Open PvP worlds.
	WorldTypeOpenPvP = WorldType{0}

	// WorldTypeOptionalPvP represents Optional PvP worlds.
	WorldTypeOptionalPvP = WorldType{1}

	// WorldTypeHardcorePvP represents Hardcore PvP worlds.
	WorldTypeHardcorePvP = WorldType{2}

	// WorldTypeRetroOpenPvP represents Retro Open PvP worlds.
	WorldTypeRetroOpenPvP = WorldType{3}

	// WorldTypeRetroHardcorePvP represents Retro Hardcore PvP worlds.
	WorldTypeRetroHardcorePvP = WorldType{4}
)

// ID returns the integer representation of the World Type.
//
// It can be used to access the numerical representation of the World Type when
// needed.
func (wt WorldType) ID() int {
	return wt.wt
}

// QueryVal returns the query parameter value representation of the World Type.
//
// The QueryVal method returns the string representation of the World Type,
// suitable for use as a query parameter value when making requests to
// tibia.com endpoints that support filtering by a world type.
//
// Example usage:
//
//	wt1 := tibia.WorldTypeOpenPvP
//	vals := url.Values{}
//	vals.Set(wt1.QueryKey(), wt1.QueryVal())
//
//	wt2 := WorldTypeRetroHardcorePvP
//	vals.Add(wt2.QueryKey(), wt2.QueryVal())
func (wt WorldType) QueryVal() string {
	return strconv.FormatInt(int64(wt.wt), 10)
}

// QueryKey returns the query parameter key for filtering by World Type.
//
// The QueryKey method returns the string representation of the query parameter
// key to be used when filtering by World Type in tibia.com requests.
// This key can be appended to the query string to specify the desired
// World Type.
//
// Example usage:
//
//	wt1 := tibia.WorldTypeOpenPvP
//	vals := url.Values{}
//	vals.Set(wt1.QueryKey(), wt1.QueryVal())
//
//	wt2 := WorldTypeRetroHardcorePvP
//	vals.Add(wt2.QueryKey(), wt2.QueryVal())
func (wt WorldType) QueryKey() string {
	return "worldtypes[]"
}

// String returns the string representation of the World Type.
func (wt WorldType) String() string {
	switch wt {
	case WorldTypeOpenPvP:
		return "Open PvP"
	case WorldTypeOptionalPvP:
		return "Optional PvP"
	case WorldTypeHardcorePvP:
		return "Hardcore PvP"
	case WorldTypeRetroOpenPvP:
		return "Retro Open PvP"
	case WorldTypeRetroHardcorePvP:
		return "Retro Hardcore PvP"
	default:
		panic("unknown wt")
	}
}
