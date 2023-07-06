package tibia

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestHighscoreJsonMarshal(t *testing.T) {
	type Test struct {
		HS HighscoreCategory `json:"highscore"`
	}

	for _, tc := range []struct {
		name  string
		input Test
		want  []byte
	}{
		{
			name:  "Achievements",
			input: Test{HighscoreCategoryAchievements},
			want:  []byte(`{"highscore":"Achievements"}`),
		},
		{
			name:  "Axe Fighting",
			input: Test{HighscoreCategoryAxeFighting},
			want:  []byte(`{"highscore":"Axe Fighting"}`),
		},
		{
			name:  "Charm Points",
			input: Test{HighscoreCategoryCharmPoints},
			want:  []byte(`{"highscore":"Charm Points"}`),
		},
		{
			name:  "Club Fighting",
			input: Test{HighscoreCategoryClubFighting},
			want:  []byte(`{"highscore":"Club Fighting"}`),
		},
		{
			name:  "Distance Fighting",
			input: Test{HighscoreCategoryDistanceFighting},
			want:  []byte(`{"highscore":"Distance Fighting"}`),
		},
		{
			name:  "Experience Points",
			input: Test{HighscoreCategoryExperiencePoints},
			want:  []byte(`{"highscore":"Experience Points"}`),
		},
		{
			name:  "Fishing",
			input: Test{HighscoreCategoryFishing},
			want:  []byte(`{"highscore":"Fishing"}`),
		},
		{
			name:  "Fist Fighting",
			input: Test{HighscoreCategoryFistFighting},
			want:  []byte(`{"highscore":"Fist Fighting"}`),
		},
		{
			name:  "Goshnars Taint",
			input: Test{HighscoreCategoryGoshnarsTaint},
			want:  []byte(`{"highscore":"Goshnar's Taint"}`),
		},
		{
			name:  "Loyalty Points",
			input: Test{HighscoreCategoryLoyaltyPoints},
			want:  []byte(`{"highscore":"Loyalty Points"}`),
		},
		{
			name:  "Magic Level",
			input: Test{HighscoreCategoryMagicLevel},
			want:  []byte(`{"highscore":"Magic Level"}`),
		},
		{
			name:  "Shielding",
			input: Test{HighscoreCategoryShielding},
			want:  []byte(`{"highscore":"Shielding"}`),
		},
		{
			name:  "Sword Fighting",
			input: Test{HighscoreCategorySwordFighting},
			want:  []byte(`{"highscore":"Sword Fighting"}`),
		},
		{
			name:  "Drome Score",
			input: Test{HighscoreCategoryDromeScore},
			want:  []byte(`{"highscore":"Drome Score"}`),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.input)
			if err != nil {
				t.Errorf("failed to marshal json: %s", err)
				return
			}

			if !bytes.Equal(tc.want, data) {
				t.Errorf(
					"unexpected marshal result\nwant: %s\ngot: %s\n",
					string(tc.want), string(data),
				)
				return
			}
		})
	}
}

func TestHighscoreJsonUnmarshal(t *testing.T) {
	type Test struct {
		HS HighscoreCategory `json:"highscore"`
	}

	for _, tc := range []struct {
		name  string
		input []byte
		want  Test
	}{
		{
			name:  "shielding",
			want:  Test{HighscoreCategoryShielding},
			input: []byte(`{"highscore":"Shielding"}`),
		},
		{
			name:  "shielding int str",
			want:  Test{HighscoreCategoryShielding},
			input: []byte(`{"highscore":"12"}`),
		},
		{
			name:  "shielding int",
			want:  Test{HighscoreCategoryShielding},
			input: []byte(`{"highscore":12}`),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var be Test
			if err := json.Unmarshal(tc.input, &be); err != nil {
				t.Errorf("failed to unmarshal json: %s", err)
				return
			}

			if be != tc.want {
				t.Errorf(
					"unexpected unmarshal result\nwant: %#v\ngot: %#v\n",
					tc.want, be,
				)
				return
			}
		})
	}
}
