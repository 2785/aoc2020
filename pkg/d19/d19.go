package d19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/2785/aoc2020/pkg/regutil"
	"golang.org/x/sync/errgroup"
)

// Rules ..
type Rules map[int]func(rules Rules) (string, error)

// ParseInput ..
func ParseInput(f []byte) (Rules, []string, error) {
	split := strings.Split(string(f), "\n\n")

	// process rules
	splitRuleLine := strings.Split(split[0], "\n")
	rules := make(Rules, len(splitRuleLine))
	ruleCache := make(map[int]string) // this is cheesy as shit
	cacheMu := &sync.Mutex{}
	for _, v := range splitRuleLine {

		splitRule := strings.Split(v, ": ")
		ruleNum, err := strconv.Atoi(splitRule[0])
		if err != nil {
			return nil, nil, fmt.Errorf("cannot parse number %s: %w", splitRule[0], err)
		}
		content := splitRule[1]
		if content[0] == '"' {
			letter := strings.Trim(content, "\"")
			rules[ruleNum] = func(r Rules) (string, error) {
				return letter, nil
			}
		} else {
			parts := strings.Split(content, "|")
			indices := make([][]int, len(parts))

			for ind, part := range parts {
				numbers := strings.Split(strings.TrimSpace(part), " ")
				indices[ind] = make([]int, len(numbers))
				for j, numStr := range numbers {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						return nil, nil, fmt.Errorf("cannot parse number %s: %w", numStr, err)
					}
					indices[ind][j] = num
				}
			}
			rules[ruleNum] = func(r Rules) (string, error) {
				if val, ok := ruleCache[ruleNum]; ok {
					return val, nil
				}

				locRuleSet := make([]string, len(indices))

				for ind, alt := range indices {
					out := ""
					for _, v := range alt {
						if val, ok := ruleCache[v]; ok {
							out += val
							continue
						}
						fn, ok := r[v]
						if !ok {
							return "", fmt.Errorf("unknown index %v", v)
						}
						val, err := fn(r)
						if err != nil {
							return "", err
						}

						if v != 8 && v != 11 {
							cacheMu.Lock()
							ruleCache[v] = val
							cacheMu.Unlock()
						}
						out += val
					}
					locRuleSet[ind] = out
				}

				return fmt.Sprintf("(?:%s)", strings.Join(locRuleSet, "|")), nil
			}
		}
	}

	// Process input

	lines := strings.Split(split[1], "\n")

	return rules, lines, nil
}

// SolvePart1 ..
func SolvePart1(rules Rules, inputs []string) (int, error) {
	reStr, err := rules[0](rules)
	if err != nil {
		return 0, err
	}
	re, err := regexp.Compile(fmt.Sprintf("^%s$", reStr))
	if err != nil {
		return 0, err
	}
	var count int64
	eg := &errgroup.Group{}
	for _, v := range inputs {
		eg.Go(func(curr string) func() error {
			return func() error {
				if re.MatchString(curr) {
					atomic.AddInt64(&count, 1)
				}
				return nil
			}
		}(v))
	}

	err = eg.Wait()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// SolvePart2 ..
func SolvePart2(rules Rules, inputs []string) (int, error) {

	p42, err := rules[42](rules)
	if err != nil {
		return 0, err
	}

	p31, err := rules[31](rules)
	if err != nil {
		return 0, err
	}

	re, err := regexp.Compile(fmt.Sprintf("^(?P<r42>%s+)(?P<r31>%s+)$", p42, p31))
	if err != nil {
		return 0, err
	}
	var count int64
	eg := &errgroup.Group{}
	for _, v := range inputs {
		eg.Go(func(curr string) func() error {
			return func() error {
				matches, err := regutil.MustCaptureNamedGroup(re, []string{"r42", "r31"}, curr)
				if err == nil {
					c42 := len(regexp.MustCompile(p42).FindAllString(matches["r42"], -1))
					c31 := len(regexp.MustCompile(p31).FindAllString(matches["r31"], -1))

					if c42 > c31 {
						atomic.AddInt64(&count, 1)
					}

				}
				return nil
			}
		}(v))
	}

	err = eg.Wait()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
