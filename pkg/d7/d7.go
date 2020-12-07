package d7

import (
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/2785/aoc2020/pkg/regutil"
	"golang.org/x/sync/errgroup"
)

// RuleBook ..
type RuleBook map[string]map[string]int

// DecodeOne ..
func DecodeOne(s string) (name string, rule map[string]int, err error) {
	firstSplit := strings.SplitN(s, "contain", 2)

	firstMatch, err := regutil.MustCaptureNamedGroup(regexp.MustCompile(`(?P<color>\w+ \w+) bags?`), []string{"color"}, firstSplit[0])

	if err != nil {
		return "", nil, err
	}

	name = firstMatch["color"]

	if regexp.MustCompile(`no other bags`).MatchString(firstSplit[1]) {
		return
	}

	secondSplit := strings.Split(firstSplit[1], ", ")

	rule = make(map[string]int)

	for _, v := range secondSplit {
		secondMatch, err := regutil.MustCaptureNamedGroup(regexp.MustCompile(`(?P<number>\d+) (?P<color>\w+ \w+) bags?`), []string{"number", "color"}, v)
		if err != nil {
			return "", nil, err
		}

		count, err := strconv.Atoi(secondMatch["number"])

		if err != nil {
			return "", nil, err
		}

		rule[secondMatch["color"]] = count
	}

	return
}

// ParseInput ..
func ParseInput(f []byte) (RuleBook, error) {
	lines := strings.Split(string(f), "\n")
	book := make(RuleBook)
	errGrp := &errgroup.Group{}
	mu := &sync.Mutex{}

	for _, line := range lines {
		errGrp.Go(func(s string) func() error {
			return func() error {
				name, rule, err := DecodeOne(s)
				if err != nil {
					return err
				}
				mu.Lock()
				book[name] = rule
				mu.Unlock()
				return nil
			}

		}(line))
	}

	err := errGrp.Wait()

	if err != nil {
		return nil, err
	}

	return book, nil
}

// SolvePart1 ..
func SolvePart1(rb RuleBook) int {
	done := false

	currSet := map[string]struct{}{"shiny gold": {}}
	fullSet := make(map[string]struct{})
	for !done {
		newSet := make(map[string]struct{})
		mu := &sync.Mutex{}
		wg := &sync.WaitGroup{}
		wg.Add(len(currSet))
		for curr := range currSet {
			go func(color string) {
				defer wg.Done()
				for out, in := range rb {
					if _, ok := in[color]; ok {
						mu.Lock()
						newSet[out] = struct{}{}
						mu.Unlock()
					}
				}
			}(curr)
		}
		wg.Wait()

		replacementSet := make(map[string]struct{})

		for k, v := range newSet {
			if _, ok := fullSet[k]; !ok {
				replacementSet[k] = v
			}
		}

		if len(replacementSet) == 0 {
			done = true
		}

		for k, v := range currSet {
			fullSet[k] = v
		}

		currSet = replacementSet
	}

	return len(fullSet) - 1
}

// SolvePart2 ..
func SolvePart2(rb RuleBook) int {
	currSet := map[string]int{"shiny gold": 1}
	fullSet := make(map[string]int)
	for len(currSet) > 0 {
		newSet := make(map[string]int)
		mu := &sync.Mutex{}
		wg := &sync.WaitGroup{}
		wg.Add(len(currSet))
		for curr, count := range currSet {
			go func(color string, number int) {
				defer wg.Done()
				for c, n := range rb[color] {
					mu.Lock()
					if _, ok := newSet[c]; !ok {
						newSet[c] = 0
					}
					newSet[c] += number * n
					mu.Unlock()
				}
			}(curr, count)
		}
		wg.Wait()

		for k, v := range newSet {
			if _, ok := fullSet[k]; !ok {
				fullSet[k] = 0
			}

			fullSet[k] += v
		}

		currSet = newSet
	}

	sum := 0

	for _, v := range fullSet {
		sum += v
	}

	return sum

}
