package d21

import (
	"regexp"
	"sort"
	"strings"

	"github.com/2785/aoc2020/pkg/regutil"
	"golang.org/x/sync/errgroup"
)

// Entry ..
type Entry struct {
	ing   []string
	aller []string
}

// ParseInput ..
func ParseInput(f []byte) ([]*Entry, error) {
	file := strings.ReplaceAll(string(f), "\r", "")
	split := strings.Split(file, "\n")

	re := regexp.MustCompile(`(?P<ing>[\w\s]+) \(contains (?P<aller>[\w\s,]+)\)`)

	eg := &errgroup.Group{}
	out := make([]*Entry, len(split))

	for i, v := range split {
		eg.Go(func(ind int, line string) func() error {
			return func() error {
				matches, err := regutil.MustCaptureNamedGroup(re, []string{"ing", "aller"}, line)
				if err != nil {
					return err
				}
				out[ind] = &Entry{
					ing:   strings.Split(matches["ing"], " "),
					aller: strings.Split(matches["aller"], ", "),
				}
				return nil
			}
		}(i, v))
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return out, nil
}

// SolvePart1 ..
func SolvePart1(records []*Entry) (int, error) {
	record := make(map[string]map[string]struct{})

	allIng := make(map[string]struct{})

	for _, v := range records {

		for _, ing := range v.ing {
			if _, ok := allIng[ing]; !ok {
				allIng[ing] = struct{}{}
			}
		}

		for _, aller := range v.aller {
			if _, ok := record[aller]; !ok {
				record[aller] = make(map[string]struct{})
				for _, ing := range v.ing {
					record[aller][ing] = struct{}{}
				}
			} else {
				for a := range record[aller] {
					if !isIn(a, v.ing) {
						delete(record[aller], a)
					}
				}
			}
		}
	}

	for func() bool {
		for _, v := range record {
			if len(v) > 1 {
				return true
			}

		}
		return false
	}() {
		for thing, v := range record {
			if len(v) == 1 {
				item := func() string {
					for k := range v {
						return k
					}
					return ""
				}()

				for thingy, rest := range record {
					if thingy != thing {
						delete(rest, item)
					}
				}
			}
		}
	}

	doesNotContainAllergen := make([]string, 0)

	for v := range allIng {
		hasAllergen := false
		for _, val := range record {
			for i := range val {
				if i == v {
					hasAllergen = true
				}
			}
		}
		if !hasAllergen {
			doesNotContainAllergen = append(doesNotContainAllergen, v)
		}
	}

	count := 0

	for _, ing := range doesNotContainAllergen {
		for _, entry := range records {
			if isIn(ing, entry.ing) {
				count++
			}
		}
	}

	return count, nil
}

// SolvePart2 ..
func SolvePart2(records []*Entry) (string, error) {
	record := make(map[string]map[string]struct{})

	allIng := make(map[string]struct{})

	for _, v := range records {

		for _, ing := range v.ing {
			if _, ok := allIng[ing]; !ok {
				allIng[ing] = struct{}{}
			}
		}

		for _, aller := range v.aller {
			if _, ok := record[aller]; !ok {
				record[aller] = make(map[string]struct{})
				for _, ing := range v.ing {
					record[aller][ing] = struct{}{}
				}
			} else {
				for a := range record[aller] {
					if !isIn(a, v.ing) {
						delete(record[aller], a)
					}
				}
			}
		}
	}

	for func() bool {
		for _, v := range record {
			if len(v) > 1 {
				return true
			}

		}
		return false
	}() {
		for thing, v := range record {
			if len(v) == 1 {
				item := func() string {
					for k := range v {
						return k
					}
					return ""
				}()

				for thingy, rest := range record {
					if thingy != thing {
						delete(rest, item)
					}
				}
			}
		}
	}

	allergens := func() []string {
		out := make([]string, 0)
		for name := range record {
			out = append(out, name)
		}
		return out
	}()

	sort.Strings(allergens)

	list := func() []string {
		out := make([]string, len(allergens))
		for i, v := range allergens {
			out[i] = func() string {
				for n := range record[v] {
					return n
				}
				panic("this should not happen")
			}()
		}
		return out
	}()

	return strings.Join(list, ","), nil
}

func isIn(s string, coll []string) bool {
	for _, v := range coll {
		if s == v {
			return true
		}
	}
	return false
}
