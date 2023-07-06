package tibia

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBEJsonMarshal(t *testing.T) {
	type Test struct {
		BE BattleEyeStatus `json:"battle_eye"`
	}

	for _, tc := range []struct {
		name  string
		input Test
		want  []byte
	}{
		{
			name:  "any world",
			input: Test{BattleEyeStatusAnyWorld},
			want:  []byte(`{"battle_eye":"Any World"}`),
		},
		{
			name:  "unprotected",
			input: Test{BattleEyeStatusUnprotected},
			want:  []byte(`{"battle_eye":"Unprotected"}`),
		},
		{
			name:  "protected",
			input: Test{BattleEyeStatusProtected},
			want:  []byte(`{"battle_eye":"Protected"}`),
		},
		{
			name:  "initially protected",
			input: Test{BattleEyeStatusInitiallyProtected},
			want:  []byte(`{"battle_eye":"Initially Protected"}`),
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

func TestBEJsonUnmarshal(t *testing.T) {
	type Test struct {
		BE BattleEyeStatus `json:"battle_eye"`
	}

	for _, tc := range []struct {
		name  string
		input []byte
		want  Test
	}{
		{
			name:  "any world",
			want:  Test{BattleEyeStatusAnyWorld},
			input: []byte(`{"battle_eye":"Any World"}`),
		},
		{
			name:  "unprotected",
			want:  Test{BattleEyeStatusUnprotected},
			input: []byte(`{"battle_eye":"Unprotected"}`),
		},
		{
			name:  "protected",
			want:  Test{BattleEyeStatusProtected},
			input: []byte(`{"battle_eye":"Protected"}`),
		},
		{
			name:  "initially protected",
			want:  Test{BattleEyeStatusInitiallyProtected},
			input: []byte(`{"battle_eye":"Initially Protected"}`),
		},
		{
			name:  "empty",
			want:  Test{},
			input: []byte(`{}`),
		},
		{
			name:  "any world str int",
			want:  Test{BattleEyeStatusAnyWorld},
			input: []byte(`{"battle_eye":"-1"}`),
		},
		{
			name:  "unprotected str int",
			want:  Test{BattleEyeStatusUnprotected},
			input: []byte(`{"battle_eye":"0"}`),
		},
		{
			name:  "protected str int",
			want:  Test{BattleEyeStatusProtected},
			input: []byte(`{"battle_eye":"1"}`),
		},
		{
			name:  "initially protected str int",
			want:  Test{BattleEyeStatusInitiallyProtected},
			input: []byte(`{"battle_eye":"2"}`),
		},
		{
			name:  "any world int",
			want:  Test{BattleEyeStatusAnyWorld},
			input: []byte(`{"battle_eye":-1}`),
		},
		{
			name:  "unprotected int",
			want:  Test{BattleEyeStatusUnprotected},
			input: []byte(`{"battle_eye":0}`),
		},
		{
			name:  "protected int",
			want:  Test{BattleEyeStatusProtected},
			input: []byte(`{"battle_eye":1}`),
		},
		{
			name:  "initially protected int",
			want:  Test{BattleEyeStatusInitiallyProtected},
			input: []byte(`{"battle_eye":2}`),
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
