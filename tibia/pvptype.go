package tibia

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// PvPTypeFromString converts a string representation of a PvP Type to its
// corresponding PvPType.
//
// This conversion allows you to work with PvP types in a more convenient and
// type-safe manner.
//
// The function performs a case-insensitive comparison of the provided string
// against known PvP types. If a match is found, the corresponding PvPType
// is returned along with a nil error.
//
// If the provided string does not match any known PvP types,
// an ErrUnknownPvPType is returned.
//
// Strings representing the integer value of a PvPType (i.e. "1" for
// Opptional PvP) will also be parsed into their corresponding PvPType.
func PvPTypeFromString(be string) (PvPType, error) {
	switch strings.ToLower(be) {
	case "open", "openpvp", "open pvp", "0":
		return PvPTypeOpenPvP, nil
	case "optional", "optionalpvp", "optional pvp", "1":
		return PvPTypeOptionalPvP, nil
	case "hardcore", "hardcorepvp", "hardcore pvp", "2":
		return PvPTypeHardcorePvP, nil
	case "retro open", "retroopen", "retro open pvp", "retroopenpvp", "3":
		return PvPTypeRetroOpenPvP, nil
	case "retro hardcore", "retrohardcore", "retro hardcore pvp",
		"retrohardcorepvp", "4":
		return PvPTypeRetroHardcorePvP, nil
	default:
		return PvPType{}, ErrUnknownPvPType
	}
}

// PvPTypeFromInt converts an integer representation of a PvP Type to its
// corresponding PvP Type.
//
// This conversion allows you to work with PvP Type in a more convenient and
// type-safe manner.
//
// The function performs a comparison of the provided integer against known
// PvP types. If a match is found, the corresponding PvP Type is returned
// along with a nil error.
//
// If the provided integer does not match any known PvP types,
// an ErrUnknownPvPType is returned.
func PvPTypeFromInt(be int) (PvPType, error) {
	switch be {
	case 0:
		return PvPTypeOpenPvP, nil
	case 1:
		return PvPTypeOptionalPvP, nil
	case 2:
		return PvPTypeHardcorePvP, nil
	case 3:
		return PvPTypeRetroOpenPvP, nil
	case 4:
		return PvPTypeRetroHardcorePvP, nil
	default:
		return PvPType{}, ErrUnknownPvPType
	}
}

// PvPType represents a tibia PvP type, such as Open PvP.
type PvPType struct {
	pt int
}

var (
	// PvPTypeOpenPvP represents Open PvP worlds.
	PvPTypeOpenPvP = PvPType{0}

	// PvPTypeOptionalPvP represents Optional PvP worlds.
	PvPTypeOptionalPvP = PvPType{1}

	// PvPTypeHardcorePvP represents Hardcore PvP worlds.
	PvPTypeHardcorePvP = PvPType{2}

	// PvPTypeRetroOpenPvP represents Retro Open PvP worlds.
	PvPTypeRetroOpenPvP = PvPType{3}

	// PvPTypeRetroHardcorePvP represents Retro Hardcore PvP worlds.
	PvPTypeRetroHardcorePvP = PvPType{4}
)

// ID returns the integer representation of the PvP Type.
//
// It can be used to access the numerical representation of the PvP Type when
// needed.
func (pt PvPType) ID() int {
	return pt.pt
}

// QueryVal returns the query parameter value representation of the PvP Type.
//
// The QueryVal method returns the string representation of the PvP Type,
// suitable for use as a query parameter value when making requests to
// tibia.com endpoints that support filtering by a PvP type.
//
// Example usage:
//
//	wt1 := tibia.PvPTypeOpenPvP
//	vals := url.Values{}
//	vals.Set(wt1.QueryKey(), wt1.QueryVal())
//
//	wt2 := PvPTypeRetroHardcorePvP
//	vals.Add(wt2.QueryKey(), wt2.QueryVal())
func (pt PvPType) QueryVal() string {
	return strconv.FormatInt(int64(pt.pt), 10)
}

// QueryKey returns the query parameter key for filtering by PvP Type.
//
// The QueryKey method returns the string representation of the query parameter
// key to be used when filtering by PvP Type in tibia.com requests.
// This key can be appended to the query string to specify the desired
// PvP Type.
//
// Example usage:
//
//	wt1 := tibia.PvPTypeOpenPvP
//	vals := url.Values{}
//	vals.Set(wt1.QueryKey(), wt1.QueryVal())
//
//	wt2 := PvPTypeRetroHardcorePvP
//	vals.Add(wt2.QueryKey(), wt2.QueryVal())
func (pt PvPType) QueryKey() string {
	return "PvPTypes[]"
}

// String returns the string representation of the PvP Type.
func (pt PvPType) String() string {
	switch pt {
	case PvPTypeOpenPvP:
		return "Open PvP"
	case PvPTypeOptionalPvP:
		return "Optional PvP"
	case PvPTypeHardcorePvP:
		return "Hardcore PvP"
	case PvPTypeRetroOpenPvP:
		return "Retro Open PvP"
	case PvPTypeRetroHardcorePvP:
		return "Retro Hardcore PvP"
	default:
		panic("unknown wt")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (pt *PvPType) UnmarshalJSON(b []byte) error {
	var zero any
	if err := json.Unmarshal(b, &zero); err != nil {
		return fmt.Errorf("failed to unmarshal pvp type: %w", err)
	}

	switch v := zero.(type) {
	case string:
		return pt.unmarshalFromString(v)
	case float64:
		return pt.unmarshalFromInt(int(v))
	default:
		return fmt.Errorf("can not unmarshal %T into pvp type", v)
	}
}

func (pt *PvPType) unmarshalFromString(data string) error {
	if data == "" {
		return nil
	}

	_pt, err := PvPTypeFromString(data)
	if err != nil {
		return fmt.Errorf("pvp type unmarshal: %w", err)
	}

	*pt = _pt
	return nil
}

func (pt *PvPType) unmarshalFromInt(data int) error {
	_pt, err := PvPTypeFromInt(data)
	if err != nil {
		return fmt.Errorf("pvp type unmarshal: %w", err)
	}

	*pt = _pt
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (pt PvPType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + pt.String() + `"`), nil
}
