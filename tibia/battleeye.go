package tibia

import (
	"strconv"
	"strings"
)

// BattleEyeStatusFromString converts a string representation of a Battle Eye
// status to its corresponding BattleEyeStatus.
//
// This conversion allows you to work with Battle Eye statuses in a more
// convenient and type-safe manner.
//
// The function performs a case-insensitive comparison of the provided string
// against known Battle Eye status values. If a match is found, the
// corresponding BattleEyeStatus is returned along with a nil error.
//
// If the provided string does not match any known Battle Eye status values,
// an ErrUnknownBattleEyeStatus is returned.
//
// Strings representing the integer value of a BattleEyeStatus (i.e. "-1" for
// Any World) will also be parsed into their corresponding BattleEyeStatus.
func BattleEyeStatusFromString(be string) (BattleEyeStatus, error) {
	switch strings.ToLower(be) {
	case "any world", "any", "anyworld", "-1":
		return BattleEyeStatusAnyWorld, nil
	case "unprotected", "0":
		return BattleEyeStatusUnprotected, nil
	case "protected", "1":
		return BattleEyeStatusProtected, nil
	case "initially protected", "initiallyprotected", "2":
		return BattleEyeStatusInitiallyProtected, nil
	default:
		return BattleEyeStatus{}, ErrUnknownBattleEyeStatus
	}
}

// BattleEyeStatusFromInt converts an integer representation of a Battle Eye
// status to its corresponding BattleEyeStatus value.
//
// This conversion allows you to work with Battle Eye statuses in a more
// convenient and type-safe manner.
//
// The function performs a comparison of the provided integer against known
// Battle Eye status values. If a match is found, the corresponding
// BattleEyeStatus value is returned along with a nil error.
//
// If the provided integer does not match any known Battle Eye status values,
// an ErrUnknownBattleEyeStatus is returned.
func BattleEyeStatusFromInt(be int) (BattleEyeStatus, error) {
	switch be {
	case -1:
		return BattleEyeStatusAnyWorld, nil
	case 0:
		return BattleEyeStatusUnprotected, nil
	case 1:
		return BattleEyeStatusProtected, nil
	case 2:
		return BattleEyeStatusInitiallyProtected, nil
	default:
		return BattleEyeStatus{}, ErrUnknownBattleEyeStatus
	}
}

// BattleEyeStatus represents the status of Battle Eye protection for a Tibia
// world.
//
// The BattleEyeStatus struct is used to represent the different Battle Eye
// protection statuses for Tibia worlds. Battle Eye is an anti-cheat system
// implemented by CipSoft to combat cheating and enhance the gaming experience
// for legitimate players.
type BattleEyeStatus struct {
	be int
}

var (
	// BattleEyeStatusAnyWorld represents a world that has any of the possible
	// Battle Eye statuses.
	BattleEyeStatusAnyWorld = BattleEyeStatus{-1}

	// BattleEyeStatusUnprotected represents a world that is not protected by
	// Battle Eye.
	BattleEyeStatusUnprotected = BattleEyeStatus{0}

	// BattleEyeStatusProtected represents a world that wasn't initially
	// protected by Battle Eye but is now being protected.
	BattleEyeStatusProtected = BattleEyeStatus{1}

	// BattleEyeStatusInitiallyProtected represents a world that has been
	// protected by Battle Eye since day one.
	BattleEyeStatusInitiallyProtected = BattleEyeStatus{2}
)

// ID returns the integer representation of the Battle Eye status.
//
// It can be used to access the numerical representation of the Battle Eye
// status when needed.
func (be BattleEyeStatus) ID() int {
	return be.be
}

// QueryVal returns the query parameter value representation of the Battle Eye
// status.
//
// The QueryVal method returns the string representation of the Battle Eye
// status, suitable for use as a query parameter value when making requests to
// tibia.com endpoints that support filtering by a Battle Eye status.
//
// Example usage:
//
//	be := tibia.BattleEyeStatusUnprotected
//	var vals url.Values
//	vals.Set(be.QueryKey(), be.QueryVal())
func (be BattleEyeStatus) QueryVal() string {
	return strconv.FormatInt(int64(be.be), 10)
}

// QueryKey returns the query parameter key for filtering by Battle Eye status.
//
// The QueryKey method returns the string representation of the query parameter
// key to be used when filtering by Battle Eye status in tibia.com requests.
// This key can be appended to the query string to specify the desired
// Battle Eye status.
//
// Example usage:
//
//	be := tibia.BattleEyeStatusUnprotected
//	var vals url.Values
//	vals.Set(be.QueryKey(), be.QueryVal())
func (be BattleEyeStatus) QueryKey() string {
	return "beprotection"
}

// String returns the string representation of the Battle Eye status.
func (be BattleEyeStatus) String() string {
	switch be {
	case BattleEyeStatusAnyWorld:
		return "Any World"
	case BattleEyeStatusUnprotected:
		return "Unprotected"
	case BattleEyeStatusProtected:
		return "Protected"
	case BattleEyeStatusInitiallyProtected:
		return "Initially Protected"
	default:
		panic("unknown be")
	}
}
