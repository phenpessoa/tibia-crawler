package tibia

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestVocationJsonMarshal(t *testing.T) {
	type Test struct {
		V Vocation `json:"vocation"`
	}

	for _, tc := range []struct {
		name  string
		input Test
		want  []byte
	}{
		{
			name:  "All",
			input: Test{VocationAll},
			want:  []byte(`{"vocation":"All"}`),
		},
		{
			name:  "None",
			input: Test{VocationNone},
			want:  []byte(`{"vocation":"None"}`),
		},
		{
			name:  "Knight",
			input: Test{VocationKnight},
			want:  []byte(`{"vocation":"Knight"}`),
		},
		{
			name:  "Paladin",
			input: Test{VocationPaladin},
			want:  []byte(`{"vocation":"Paladin"}`),
		},
		{
			name:  "Sorcerer",
			input: Test{VocationSorcerer},
			want:  []byte(`{"vocation":"Sorcerer"}`),
		},
		{
			name:  "Druid",
			input: Test{VocationDruid},
			want:  []byte(`{"vocation":"Druid"}`),
		},
		{
			name:  "Elite Knight",
			input: Test{VocationEliteKnight},
			want:  []byte(`{"vocation":"Elite Knight"}`),
		},
		{
			name:  "Royal Paladin",
			input: Test{VocationRoyalPaladin},
			want:  []byte(`{"vocation":"Royal Paladin"}`),
		},
		{
			name:  "Master Sorcerer",
			input: Test{VocationMasterSorcerer},
			want:  []byte(`{"vocation":"Master Sorcerer"}`),
		},
		{
			name:  "Elder Druid",
			input: Test{VocationElderDruid},
			want:  []byte(`{"vocation":"Elder Druid"}`),
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

func TestVocationJsonUnmarshal(t *testing.T) {
	type Test struct {
		V Vocation `json:"vocation"`
	}

	for _, tc := range []struct {
		name  string
		input []byte
		want  Test
	}{
		{
			name:  "sorcerer",
			want:  Test{VocationSorcerer},
			input: []byte(`{"vocation":"sorcerer"}`),
		},
		{
			name:  "sorcerer int str",
			want:  Test{VocationSorcerer},
			input: []byte(`{"vocation":"4"}`),
		},
		{
			name:  "sorcerer int",
			want:  Test{VocationSorcerer},
			input: []byte(`{"vocation":4}`),
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
