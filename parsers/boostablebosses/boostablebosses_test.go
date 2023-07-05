package boostablebosses

import (
	"io"
	"strings"
	"testing"

	"github.com/phenpessoa/tibia-crawler/internal/static"
	"github.com/phenpessoa/tibia-crawler/tibia"
)

func TestParser(t *testing.T) {
	f, err := static.TestData.Open("testdata/boostablebosses.html")
	if err != nil {
		t.Errorf("failed to open test data: %s\n%#v\n", err, err)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file: %s", err)
	}

	p := Parser{}

	if err := p.parse(string(data)); err != nil {
		t.Errorf("failed to parse data: %s\n%#v\n", err, err)
		return
	}

	boosted := p.cachedBosses.Boosted
	bosses := p.cachedBosses.Bosses

	if len(bosses) != tibia.AmountOfBoostableBosses {
		t.Errorf(
			"Wrong length\nwant: %d\ngot: %d",
			tibia.AmountOfBoostableBosses, len(bosses),
		)
		return
	}

	if boosted.Name != "Utua Stone Sting" {
		t.Errorf(
			"Wrong boosted\nwant: %s\ngot: %s",
			"Utua Stone Sting", boosted.Name,
		)
		return
	}

	bi := "https://static.tibia.com/images/global/header/monsters/utua.gif"
	if boosted.ImageURL != bi {
		t.Errorf(
			"Wrong boosted image\nwant: %s\ngot: %s",
			bi, boosted.ImageURL,
		)
		return
	}

	for _, tc := range []struct {
		idx       int
		name      string
		isBoosted bool
		imageURL  string
	}{
		{
			idx:       19,
			name:      "Gnomevil",
			isBoosted: false,
			imageURL:  "gnomehorticulist.gif",
		},
		{
			idx:       24,
			name:      "Goshnar's Malice",
			isBoosted: false,
			imageURL:  "goshnarsmalice.gif",
		},
		{
			idx:       52,
			name:      "Sharpclaw",
			isBoosted: false,
			imageURL:  "sharpclaw.gif",
		},
		{
			idx:       75,
			name:      "The Pale Worm",
			isBoosted: false,
			imageURL:  "paleworm.gif",
		},
		{
			idx:       87,
			name:      "Utua Stone Sting",
			isBoosted: true,
			imageURL:  "utua.gif",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			boss := bosses[tc.idx]

			if tc.name != boss.Name {
				t.Errorf(
					"Wrong name\nidx: %d (%s)\nwant: %s\ngot: %s",
					tc.idx, tc.name, tc.name, boss.Name,
				)
			}

			if tc.isBoosted != boss.IsBoosted {
				t.Errorf(
					"Wrong isBoosted status\nidx: %d (%s)\nwant: %v\ngot: %v",
					tc.idx, tc.name, tc.isBoosted, boss.IsBoosted,
				)
			}

			if !strings.Contains(boss.ImageURL, tc.imageURL) {
				t.Errorf(
					"Wrong image URL\nidx: %d (%s)\nwant: %s\ngot: %s",
					tc.idx, tc.name, tc.imageURL, boss.ImageURL,
				)
			}
		})
	}
}
