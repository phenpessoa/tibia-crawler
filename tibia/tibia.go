// Package tibia provides constants, utility functions, and models related to
// the Tibia game.
//
// These utilities and models can be used in applications that require
// information or calculations specific to the Tibia game.
package tibia

import "errors"

const (
	// AmountOfBoostableBosses is the amount of boostable bosses.
	AmountOfBoostableBosses = 91

	// MaxRunesAllowedInCharName is the maximum amount of runes allowed in a
	// character name.
	//
	// This value is the same for both new and legacy names.
	MaxRunesAllowedInCharName = 29

	// MinRunesAllowedInCharName is the minimum amount of runes allowed in a
	// character name.
	//
	// This value is the same for both new and legacy names.
	MinRunesAllowedInCharName = 2

	// MaxRunesAllowedInCharWord is the maximum amount of runes allowed in a
	// character word.
	//
	// This limit is smaller for new names. Use MaxRunesAllowedInCharWordNew
	// instead to validate new character names.
	MaxRunesAllowedInCharWord = 16

	// MaxRunesAllowedInCharWordNew is the maximum amount of runes allowed in a
	// character word.
	//
	// This limit is greater for legacy names. Use MaxRunesAllowedInCharWord
	// instead to validate legacy character names.
	MaxRunesAllowedInCharWordNew = 14

	// MinRunesAllowedInCharWord is the minimum amount of runes allowed in a
	// character word.
	//
	// This value is the same for both new and legacy names.
	MinRunesAllowedInCharWord = 2
)

var (
	// ErrUnknownBattleEyeStatus will be used when an uknown battle eye status
	// was tried to be parsed.
	ErrUnknownBattleEyeStatus = errors.New("unknown battle eye")

	// ErrUnknownVocation will be used when an uknown vocation was tried to be
	// parsed.
	ErrUnknownVocation = errors.New("unknown vocation")

	// ErrUnknownHighscoreCategory will be used when an uknown highscore
	// category was tried to be parsed.
	ErrUnknownHighscoreCategory = errors.New("unknown highscore category")

	// ErrUnknownPvPType will be used when an uknown PvP type was tried to
	// be parsed.
	ErrUnknownPvPType = errors.New("unknown pvp type")
)
