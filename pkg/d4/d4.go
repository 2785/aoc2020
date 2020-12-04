package d4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	byr string = "byr"
	iyr string = "iyr"
	eyr string = "eyr"
	hgt string = "hgt"
	hcl string = "hcl"
	ecl string = "ecl"
	pid string = "pid"
	cid string = "cid"
)

// Document ..
type Document map[string]string

// Valid ..
func (d Document) Valid() bool {
	for _, v := range []string{byr, iyr, eyr, hgt, hcl, ecl, pid} {
		if _, ok := d[v]; !ok {
			return false
		}
	}
	return true
}

// ValidStrict ...
func (d Document) ValidStrict() bool {
	if !d.Valid() {
		return false
	}

	yearRe := regexp.MustCompile(`\d{4}`)

	// year validations
	birthYear, err := strconv.Atoi(d[byr])
	if err != nil || !yearRe.MatchString(d[byr]) || !(birthYear >= 1920 && birthYear <= 2002) {
		return false
	}
	issueYear, err := strconv.Atoi(d[iyr])
	if err != nil || !yearRe.MatchString(d[iyr]) || !(issueYear >= 2010 && issueYear <= 2020) {
		return false
	}
	expYear, err := strconv.Atoi(d[eyr])
	if err != nil || !yearRe.MatchString(d[eyr]) || !(expYear >= 2020 && expYear <= 2030) {
		return false
	}

	// height validation

	height := d[hgt]
	hgtRe := regexp.MustCompile(`(?P<val>\d+)(?P<unit>cm|in)`)
	hgtMatches := hgtRe.FindStringSubmatch(height)
	if hgtMatches == nil {
		return false
	}
	hgtResult := make(map[string]string)
	for ind, name := range hgtRe.SubexpNames() {
		if ind != 0 && name != "" {
			hgtResult[name] = hgtMatches[ind]
		}
	}
	val, ok := hgtResult["val"]
	if !ok {
		return false
	}
	hgtNum, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	switch hgtResult["unit"] {
	case "in":
		if !(hgtNum >= 59 && hgtNum <= 76) {
			return false
		}
	case "cm":
		if !(hgtNum >= 150 && hgtNum <= 193) {
			return false
		}
	default:
		return false
	}

	// hcl validation
	hclRe := regexp.MustCompile(`^\#[0-9a-f]{6}$`)
	if !hclRe.MatchString(d[hcl]) {
		return false
	}

	// ecl validation
	eclRe := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	if !eclRe.MatchString(d[ecl]) {
		return false
	}

	pidRe := regexp.MustCompile(`^\d{9}$`)
	if !pidRe.MatchString(d[pid]) {
		return false
	}

	return true
}

// ParseInput ...
func ParseInput(f []byte) ([]Document, error) {
	docStrings := strings.Split(string(f), "\n\n")
	docs := make([]Document, len(docStrings))
	for i, v := range docStrings {
		spaceDel := strings.ReplaceAll(v, "\n", " ")
		split := strings.Split(spaceDel, " ")
		docs[i] = make(Document)
		for _, kvp := range split {
			kvpSplit := strings.Split(kvp, ":")
			if len(kvpSplit) != 2 {
				return nil, fmt.Errorf("error parsing kvp %s", kvp)
			}
			docs[i][kvpSplit[0]] = kvpSplit[1]
		}
	}

	return docs, nil
}

// SolvePart1 ..
func SolvePart1(docs []Document) int {
	var counter int64
	wg := &sync.WaitGroup{}
	wg.Add(len(docs))
	for _, doc := range docs {
		go func(d Document) {
			defer wg.Done()
			if d.Valid() {
				atomic.AddInt64(&counter, 1)
			}
		}(doc)
	}

	wg.Wait()

	return int(counter)
}

// SolvePart2 ..
func SolvePart2(docs []Document) int {
	var counter int64
	wg := &sync.WaitGroup{}
	wg.Add(len(docs))
	for _, doc := range docs {
		go func(d Document) {
			defer wg.Done()
			if d.ValidStrict() {
				atomic.AddInt64(&counter, 1)
			}
		}(doc)
	}

	wg.Wait()

	return int(counter)
}
