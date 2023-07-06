package tibia

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPVPTypeJsonMarshal(t *testing.T) {
	type Test struct {
		PT PvPType `json:"pvp_type"`
	}

	for _, tc := range []struct {
		name  string
		input Test
		want  []byte
	}{
		{
			name:  "Open PvP",
			input: Test{PvPTypeOpenPvP},
			want:  []byte(`{"pvp_type":"Open PvP"}`),
		},
		{
			name:  "Optional PvP",
			input: Test{PvPTypeOptionalPvP},
			want:  []byte(`{"pvp_type":"Optional PvP"}`),
		},
		{
			name:  "Hardcore PvP",
			input: Test{PvPTypeHardcorePvP},
			want:  []byte(`{"pvp_type":"Hardcore PvP"}`),
		},
		{
			name:  "Retro Open PvP",
			input: Test{PvPTypeRetroOpenPvP},
			want:  []byte(`{"pvp_type":"Retro Open PvP"}`),
		},
		{
			name:  "Retro Hardcore PvP",
			input: Test{PvPTypeRetroHardcorePvP},
			want:  []byte(`{"pvp_type":"Retro Hardcore PvP"}`),
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

func TestPVPTypeJsonUnmarshal(t *testing.T) {
	type Test struct {
		PT PvPType `json:"pvp_type"`
	}

	for _, tc := range []struct {
		name  string
		input []byte
		want  Test
	}{
		{
			name:  "open pvp",
			want:  Test{PvPTypeOpenPvP},
			input: []byte(`{"pvp_type":"open pvp"}`),
		},
		{
			name:  "open pvp int str",
			want:  Test{PvPTypeOpenPvP},
			input: []byte(`{"pvp_type":"0"}`),
		},
		{
			name:  "open pvp int",
			want:  Test{PvPTypeOpenPvP},
			input: []byte(`{"pvp_type":0}`),
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
