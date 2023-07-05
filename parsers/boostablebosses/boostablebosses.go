// Package boostablebosses provides an implementation of the Parser interface
// for parsing information about boostable bosses from the tibia.com Boostable
// Bosses Library page.
//
// To use the boostablebosses package, create an instance of the
// Parser struct, which implements the Parser interface.
// The Parse method can then be called to fetch the HTML content from the
// Boostable Bosses Library page, parse it, and return the parsed data.
// Additionally, the URL method can be used to retrieve the specific tibia.com
// endpoint being parsed.
package boostablebosses

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/phenpessoa/tibia-crawler/parsers"
	"github.com/phenpessoa/tibia-crawler/tibia"
)

const (
	tibiaServerSaveTimeLayout   = "15:04"
	tibiaServerSaveStartTimeStr = "07:58"
	tibiaServerSaveEndTimeStr   = "09:00"

	endpoint = "/library/?subtopic=boostablebosses"

	// contentLength is the aprox Content-Length of the data returned by
	// the bosstable bosses endpoint.
	contentLength = 110000
)

var (
	tibiaServerSaveStartTime, _ = time.Parse(
		tibiaServerSaveTimeLayout, tibiaServerSaveStartTimeStr,
	)

	tibiaServerSaveEndTime, _ = time.Parse(
		tibiaServerSaveTimeLayout, tibiaServerSaveEndTimeStr,
	)
)

func init() {
	tibiaServerSaveStartTime = tibiaServerSaveStartTime.UTC()
	tibiaServerSaveEndTime = tibiaServerSaveEndTime.UTC()
}

var _ parsers.Parser[Args, tibia.BoostableBosses] = (*Parser)(nil)

// Parser is an implementation of the Parser interface for
// parsing information about boostable bosses from the tibia.com Boostable
// Bosses Library page.
type Parser struct {
	mu           sync.RWMutex
	cachedBosses tibia.BoostableBosses
}

// Args is used by Parser to implement the parsers.Parser interface, but it is
// not used by this implementation.
type Args struct{}

// URL implements the parsers.Parser interface.
func (p *Parser) URL() string {
	return parsers.BaseURL + endpoint
}

// Parse implements the parsers.Parser interface.
func (p *Parser) Parse(
	ctx context.Context,
	args Args,
	opts parsers.Options,
) (tibia.BoostableBosses, error) {
	updated, err := p.maybeUpdateCache(ctx, args, opts)
	if err != nil {
		return tibia.BoostableBosses{}, err
	}

	if updated {
		return p.load(), nil
	}

	if opts.DisallowCachedResponses {
		if err := p.fetch(ctx, args, opts); err != nil {
			return tibia.BoostableBosses{}, err
		}
		return p.load(), nil
	}

	return p.load(), nil
}

func (p *Parser) maybeUpdateCache(
	ctx context.Context,
	args Args,
	opts parsers.Options,
) (bool, error) {
	if !p.isServerSaveTime() {
		return false, nil
	}

	if err := p.fetch(ctx, args, opts); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Parser) isServerSaveTime() bool {
	now := time.Now().UTC()
	return now.After(tibiaServerSaveStartTime) &&
		now.Before(tibiaServerSaveEndTime)
}

func (p *Parser) fetch(
	ctx context.Context,
	args Args,
	opts parsers.Options,
) error {
	select {
	case <-ctx.Done():
		return parsers.ErrCtxDone
	default:
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.URL(), nil)
	if err != nil {
		return fmt.Errorf("boostable bosses: failed to create req: %w", err)
	}

	if opts.RateLimiter != nil {
		opts.RateLimiter.Take()
	}

	res, err := opts.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("boostable bosses: failed to make req: %w", err)
	}
	defer res.Body.Close()

	defer io.Copy(io.Discard, res.Body)
	buf := make([]byte, contentLength)
	if _, err := io.ReadAtLeast(res.Body, buf, 1); err != nil {
		return fmt.Errorf("boostable bosses: failed to read body: %w", err)
	}

	data := unsafe.String(unsafe.SliceData(buf), len(buf))
	if err := p.parse(data); err != nil {
		return fmt.Errorf("boostable bosses: failed to parse body: %w", err)
	}

	return nil
}

const (
	startIndexer = `<div class="main-content Content">`
	endIndexer   = `<div id="Footer" class="main-footer">`

	todayChecker  = `Today's boosted boss: `
	bossesChecker = `<div class="CaptionContainer">`

	todayBossIndexer    = `title="` + todayChecker
	endTodayBossIndexer = `" src="`

	todayBossImgIndexer = `https://static.tibia.com/images/` +
		`global/header/monsters/`
	endTodayBossImgIndexer = `" onClick="`

	bossesImgIndexer    = `https://static.tibia.com/images/library/`
	endBossesImgIndexer = `"`

	bossesNameIndexer    = `border=0 /> <div>`
	endBossesNameIndexer = `</div>`
)

func (p *Parser) parse(data string) error {
	startIdx := strings.Index(
		data, startIndexer,
	)
	if startIdx == -1 {
		return fmt.Errorf("boostable bosses: main content not found")
	}

	endIdx := strings.Index(
		data[startIdx:], endIndexer,
	) + startIdx
	if endIdx == -1 {
		return fmt.Errorf("boostable bosses: end of content not found")
	}

	data = data[startIdx:endIdx]

	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("boostable bosses: no lines found")
	}

	var (
		parsed  tibia.BoostableBosses
		started bool
	)

	parsed.Boosted.IsBoosted = true
	parsed.Bosses = make(
		[]tibia.BoostableBoss, 0, tibia.AmountOfBoostableBosses,
	)

	for _, line := range lines {
		isTodaysLine := strings.Contains(line, todayChecker) && !started
		isBossesLine := strings.Contains(line, bossesChecker)

		if !isTodaysLine && !isBossesLine {
			continue
		}

		if isTodaysLine {
			started = true

			todayBossIdx := strings.Index(
				line, todayBossIndexer,
			) + len(todayBossIndexer)

			if todayBossIdx == -1 {
				return fmt.Errorf("boostable bosses: today boss idx not found")
			}

			endTodayBossIdx := strings.Index(
				line[todayBossIdx:], endTodayBossIndexer,
			) + todayBossIdx

			if endTodayBossIdx == -1 {
				return fmt.Errorf(
					"boostable bosses: today boss end idx not found",
				)
			}

			parsed.Boosted.Name = line[todayBossIdx:endTodayBossIdx]

			todayBossImgIdx := strings.Index(
				line[todayBossIdx:], todayBossImgIndexer,
			) + todayBossIdx

			if todayBossImgIdx == -1 {
				return fmt.Errorf(
					"boostable bosses: today boss img idx not found",
				)
			}

			endTodayBossImgIdx := strings.Index(
				line[todayBossImgIdx:], endTodayBossImgIndexer,
			) + todayBossImgIdx

			if endTodayBossImgIdx == -1 {
				return fmt.Errorf(
					"boostable bosses: today boss end img idx not found",
				)
			}

			parsed.Boosted.ImageURL = line[todayBossImgIdx:endTodayBossImgIdx]
		}

		if isBossesLine {
			idx := strings.Index(line, bossesImgIndexer)
			for ; idx != -1; idx = strings.Index(line, bossesImgIndexer) {
				imgIdx := strings.Index(
					line, bossesImgIndexer,
				)

				if imgIdx == -1 {
					return fmt.Errorf(
						"boostable bosses: img idx not found",
					)
				}

				endImgIdx := strings.Index(
					line[imgIdx:], endBossesImgIndexer,
				) + imgIdx

				if endImgIdx == -1 {
					return fmt.Errorf(
						"boostable bosses: end img idx not found",
					)
				}

				img := line[imgIdx:endImgIdx]

				nameIdx := strings.Index(
					line, bossesNameIndexer,
				) + len(bossesNameIndexer)

				if nameIdx == -1 {
					return fmt.Errorf(
						"boostable bosses: name idx not found",
					)
				}

				endNameIdx := strings.Index(
					line[nameIdx:], endBossesNameIndexer,
				) + nameIdx

				if endNameIdx == -1 {
					return fmt.Errorf(
						"boostable bosses: end name idx not found",
					)
				}

				name := line[nameIdx:endNameIdx]

				parsed.Bosses = append(parsed.Bosses, tibia.BoostableBoss{
					Name:      name,
					ImageURL:  img,
					IsBoosted: name == parsed.Boosted.Name,
				})

				line = line[endNameIdx-1:]
			}

			break
		}
	}

	p.store(parsed)
	return nil
}

func (p *Parser) load() tibia.BoostableBosses {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.cachedBosses
}

func (p *Parser) store(bosses tibia.BoostableBosses) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cachedBosses = bosses
}
