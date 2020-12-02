package d2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// Record ...
type Record struct {
	Letter  string
	Min     int
	Max     int
	Content string
}

var re = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<toc>\w): (?P<content>\w+)`)

// ParseInput ..
func ParseInput(f []byte) ([]Record, error) {
	split := strings.Split(string(f), "\n")
	out := make([]Record, len(split))

	for i, v := range split {
		match := re.FindStringSubmatch(v)
		result := make(map[string]string)
		for ind, name := range re.SubexpNames() {
			if ind != 0 && name != "" {
				result[name] = match[ind]
			}
		}
		for _, name := range []string{"min", "max", "toc", "content"} {
			val, ok := result[name]

			if !ok || val == "" {
				return nil, fmt.Errorf("no match found in string %v", v)
			}
		}

		min, err := strconv.Atoi(result["min"])
		if err != nil {
			return nil, fmt.Errorf("unable to process %v: %w", min, err)
		}

		max, err := strconv.Atoi(result["max"])
		if err != nil {
			return nil, fmt.Errorf("unable to process %v: %w", max, err)
		}
		out[i] = Record{
			Letter:  result["toc"],
			Content: result["content"],
			Min:     min,
			Max:     max,
		}
	}

	return out, nil
}

// SolvePart1 ..
func SolvePart1(records []Record) int {
	var counter int64
	wg := &sync.WaitGroup{}
	wg.Add(len(records))
	for _, r := range records {
		go func(curr Record) {
			numToken := strings.Count(curr.Content, curr.Letter)
			if numToken >= curr.Min && numToken <= curr.Max {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}(r)
	}

	wg.Wait()

	return int(counter)
}

// SolvePart2 ..
func SolvePart2(records []Record) int {
	var counter int64
	wg := &sync.WaitGroup{}
	wg.Add(len(records))
	for _, r := range records {
		go func(curr Record) {
			reg := regexp.MustCompile(curr.Letter)
			pos := reg.FindAllStringSubmatchIndex(curr.Content, -1)
			var count int

			for _, v := range pos {
				if v[0]+1 == curr.Min || v[0]+1 == curr.Max {
					count++
				}
			}

			if count == 1 {
				atomic.AddInt64(&counter, 1)
			}

			wg.Done()
		}(r)
	}

	wg.Wait()

	return int(counter)
}
