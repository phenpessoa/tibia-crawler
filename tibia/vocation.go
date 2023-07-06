package tibia

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// VocationFromString converts a string representation of a tibia vocation to
// its corresponding Vocation.
//
// This conversion allows you to work with vocations in a more convenient and
// type-safe manner.
//
// The function performs a case-insensitive comparison of the provided string
// against known tibia vocations. If a match is found, the corresponding
// Vocation is returned along with a nil error.
//
// If the provided string does not match any known tibia vocations,
// an ErrUnknownVocation is returned.
//
// Strings representing the integer value of a Vocation (i.e. "0" for
// All) will also be parsed into their corresponding Vocation.
func VocationFromString(vocation string) (Vocation, error) {
	switch strings.ToLower(vocation) {
	case "all", "0":
		return VocationAll, nil
	case "none", "1":
		return VocationNone, nil
	case "knight", "knights", "2":
		return VocationKnight, nil
	case "paladin", "paladins", "3":
		return VocationPaladin, nil
	case "sorcerer", "sorcerers", "4":
		return VocationSorcerer, nil
	case "druid", "druids", "5":
		return VocationDruid, nil
	case "elite knight", "elite knights", "6":
		return VocationEliteKnight, nil
	case "royal paladin", "royal paladins", "7":
		return VocationRoyalPaladin, nil
	case "master sorcerer", "master sorcerers", "8":
		return VocationMasterSorcerer, nil
	case "elder druid", "elder druids", "9":
		return VocationElderDruid, nil
	default:
		return VocationAll, ErrUnknownVocation
	}
}

// VocationFromInt converts a integer representation of a tibia vocation to
// its corresponding Vocation.
//
// This conversion allows you to work with vocations in a more convenient and
// type-safe manner.
//
// The function performs a comparison of the provided integer against known
// tibia vocations. If a match is found, the corresponding Vocation is returned
// along with a nil error.
//
// If the provided integer does not match any known tibia vocations,
// an ErrUnknownVocation is returned.
func VocationFromInt(vocation int) (Vocation, error) {
	switch vocation {
	case 0:
		return VocationAll, nil
	case 1:
		return VocationNone, nil
	case 2:
		return VocationKnight, nil
	case 3:
		return VocationPaladin, nil
	case 4:
		return VocationSorcerer, nil
	case 5:
		return VocationDruid, nil
	case 6:
		return VocationEliteKnight, nil
	case 7:
		return VocationRoyalPaladin, nil
	case 8:
		return VocationMasterSorcerer, nil
	case 9:
		return VocationElderDruid, nil
	default:
		return VocationAll, ErrUnknownVocation
	}
}

// Vocation represents a tibia vocation.
type Vocation struct {
	v int
}

var (
	// VocationAll represents the All vocation.
	VocationAll = Vocation{0}

	// VocationNone represents the None vocation.
	VocationNone = Vocation{1}

	// VocationKnight represents the Knight vocation.
	VocationKnight = Vocation{2}

	// VocationPaladin represents the Paladin vocation.
	VocationPaladin = Vocation{3}

	// VocationSorcerer represents the Sorcerer vocation.
	VocationSorcerer = Vocation{4}

	// VocationDruid represents the Druid vocation.
	VocationDruid = Vocation{5}

	// VocationEliteKnight represents the Elite Knight vocation.
	VocationEliteKnight = Vocation{6}

	// VocationRoyalPaladin represents the Royal Paladin vocation.
	VocationRoyalPaladin = Vocation{7}

	// VocationMasterSorcerer represents the Master Sorcerer vocation.
	VocationMasterSorcerer = Vocation{8}

	// VocationElderDruid represents the Elder Druid vocation.
	VocationElderDruid = Vocation{9}
)

// ID returns the integer representation of the Vocation.
//
// It can be used to access the numerical representation of the Vocation when
// needed.
//
// Note ID returns the Vocation actual ID. tibia.com does not allow filtering by
// promoted vocations on their endpoints. So, if you need to get the query ID of
// the vocation regardless of their promotion status, use the QueryID method
// instead.
func (v Vocation) ID() int {
	return v.v
}

// QueryID returns the integer representation of the Vocation, regardless of
// promotion.
//
// Note QueryID returns the Vocation query ID, not its actual ID. tibia.com does
// not allow filtering by promoted vocations on their endpoints. So, if the
// vocation is a promoted vocation, QueryID will return the same ID as the base
// vocation.
func (v Vocation) QueryID() int {
	switch v {
	case VocationAll:
		return 0
	case VocationNone:
		return 1
	case VocationKnight, VocationEliteKnight:
		return 2
	case VocationPaladin, VocationRoyalPaladin:
		return 3
	case VocationSorcerer, VocationMasterSorcerer:
		return 4
	case VocationDruid, VocationElderDruid:
		return 5
	default:
		return 0
	}
}

// QueryVal returns the query parameter value representation of the Vocation.
//
// The QueryVal method returns the string representation of the Vocation,
// suitable for use as a query parameter value when making requests to
// tibia.com endpoints that support filtering by a vocation.
//
// Note that promoted vocations and base vocations have the same QueryVal.
//
// Example usage:
//
//		v := tibia.VocationKnight
//	 // VocationEliteKnight would have the same QueryVal
//		vals := url.Values{}
//		vals.Set(v.QueryKey(), v.QueryVal())
func (v Vocation) QueryVal() string {
	return strconv.FormatInt(int64(v.QueryID()), 10)
}

// QueryKey returns the query parameter key for filtering by Vocation.
//
// The QueryKey method returns the string representation of the query parameter
// key to be used when filtering by a Vocation in tibia.com requests. This key
// can be appended to the query string to specify the desired vocation.
//
// Example usage:
//
//		v := tibia.VocationKnight
//	 // VocationEliteKnight would have the same QueryVal
//		vals := url.Values{}
//		vals.Set(v.QueryKey(), v.QueryVal())
func (v Vocation) QueryKey() string {
	return "profession"
}

// String returns the string representation of the Vocation.
func (v Vocation) String() string {
	switch v {
	case VocationAll:
		return "All"
	case VocationNone:
		return "None"
	case VocationKnight:
		return "Knight"
	case VocationEliteKnight:
		return "Elite Knight"
	case VocationPaladin:
		return "Paladin"
	case VocationRoyalPaladin:
		return "Royal Paladin"
	case VocationSorcerer:
		return "Sorcerer"
	case VocationMasterSorcerer:
		return "Master Sorcerer"
	case VocationDruid:
		return "Druid"
	case VocationElderDruid:
		return "Elder Druid"
	default:
		panic("unknown vocation")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *Vocation) UnmarshalJSON(b []byte) error {
	var zero any
	if err := json.Unmarshal(b, &zero); err != nil {
		return fmt.Errorf("failed to unmarshal vocation: %w", err)
	}

	switch vt := zero.(type) {
	case string:
		return v.unmarshalFromString(vt)
	case float64:
		return v.unmarshalFromInt(int(vt))
	default:
		return fmt.Errorf("can not unmarshal %T into vocation", vt)
	}
}

func (v *Vocation) unmarshalFromString(data string) error {
	if data == "" {
		return nil
	}

	_v, err := VocationFromString(data)
	if err != nil {
		return fmt.Errorf("vocation unmarshal: %w", err)
	}

	*v = _v
	return nil
}

func (v *Vocation) unmarshalFromInt(data int) error {
	_v, err := VocationFromInt(data)
	if err != nil {
		return fmt.Errorf("vocation unmarshal: %w", err)
	}

	*v = _v
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (v Vocation) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}
