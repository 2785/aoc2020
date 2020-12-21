package d20

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/2785/aoc2020/pkg/regutil"
	"golang.org/x/sync/errgroup"
)

type direction int

const (
	// N ..
	N direction = iota
	// E ..
	E
	// S ..
	S
	// W ..
	W
)

// NewTile ..
func NewTile(id int, data [][]bool) (*Tile, error) {
	l := len(data)
	for _, v := range data {
		if len(v) != l {
			return nil, errors.New("dimension mismatch")
		}
	}

	var n, s, e, w string
	for i := 0; i < l; i++ {
		if data[0][i] {
			n += "#"
		} else {
			n += "."
		}
		if data[l-1][i] {
			s += "#"
		} else {
			s += "."
		}
		if data[i][0] {
			w += "#"
		} else {
			w += "."
		}
		if data[i][l-1] {
			e += "#"
		} else {
			e += "."
		}
	}

	w = invertString(w)
	s = invertString(s)
	return &Tile{
		ID:   id,
		Data: data,
		Edges: map[direction]string{
			N: n,
			S: s,
			E: e,
			W: w,
		},
	}, nil
}

// Tile ..
type Tile struct {
	ID        int
	Data      [][]bool
	Edges     map[direction]string
	Neighbors map[direction]int
}

func (t *Tile) toString() string {
	out := make([]string, len(t.Data))
	for i, r := range t.Data {
		s := ""
		for _, v := range r {
			if v {
				s += "#"
			} else {
				s += "."
			}
		}
		out[i] = s
	}
	return strings.Join(out, "\n")
}

// rotateClockWise ..
func (t *Tile) rotateClockWise() *Tile {
	newData := make([][]bool, len(t.Data))
	for i := range newData {
		newData[i] = make([]bool, len(t.Data))
	}
	for i, r := range t.Data {
		for j, v := range r {
			newData[j][len(t.Data)-1-i] = v
		}
	}

	newTile := &Tile{ID: t.ID}

	newTile.Data = newData

	newTile.Edges = make(map[direction]string)

	n := t.Edges[N]
	s := t.Edges[S]
	e := t.Edges[E]
	w := t.Edges[W]

	newTile.Edges[N] = w
	newTile.Edges[E] = n
	newTile.Edges[S] = e
	newTile.Edges[W] = s

	if t.Neighbors != nil {
		newTile.Neighbors = make(map[direction]int)
		if v, ok := t.Neighbors[N]; ok {
			newTile.Neighbors[E] = v
		}
		if v, ok := t.Neighbors[E]; ok {
			newTile.Neighbors[S] = v
		}
		if v, ok := t.Neighbors[S]; ok {
			newTile.Neighbors[W] = v
		}
		if v, ok := t.Neighbors[W]; ok {
			newTile.Neighbors[N] = v
		}
	}

	return newTile
}

func (t *Tile) mirrorVertical() *Tile {
	newData := make([][]bool, len(t.Data))
	for i := range newData {
		newData[i] = make([]bool, len(t.Data))
	}
	for i, r := range t.Data {
		for j, v := range r {
			newData[i][len(t.Data)-1-j] = v
		}
	}

	newTile := &Tile{ID: t.ID}

	newTile.Data = newData
	newTile.Edges = make(map[direction]string)

	n := t.Edges[N]
	s := t.Edges[S]
	e := t.Edges[E]
	w := t.Edges[W]

	newTile.Edges[N] = invertString(n)
	newTile.Edges[E] = invertString(w)
	newTile.Edges[S] = invertString(s)
	newTile.Edges[W] = invertString(e)

	if t.Neighbors != nil {
		newTile.Neighbors = make(map[direction]int)
		if v, ok := t.Neighbors[E]; ok {
			newTile.Neighbors[W] = v
		}
		if v, ok := t.Neighbors[W]; ok {
			newTile.Neighbors[E] = v
		}
		if v, ok := t.Neighbors[N]; ok {
			newTile.Neighbors[N] = v
		}
		if v, ok := t.Neighbors[S]; ok {
			newTile.Neighbors[S] = v
		}
	}

	return newTile
}

func (t *Tile) mirrorTopDown() *Tile {
	newData := make([][]bool, len(t.Data))
	for i := range newData {
		newData[i] = make([]bool, len(t.Data))
	}
	for i, r := range t.Data {
		for j, v := range r {
			newData[len(t.Data)-1-i][j] = v
		}
	}

	newTile := &Tile{ID: t.ID}

	newTile.Data = newData
	newTile.Edges = make(map[direction]string)

	n := t.Edges[N]
	s := t.Edges[S]
	e := t.Edges[E]
	w := t.Edges[W]

	newTile.Edges[N] = invertString(s)
	newTile.Edges[E] = invertString(e)
	newTile.Edges[S] = invertString(n)
	newTile.Edges[W] = invertString(w)

	if t.Neighbors != nil {
		newTile.Neighbors = make(map[direction]int)
		if v, ok := t.Neighbors[S]; ok {
			newTile.Neighbors[N] = v
		}
		if v, ok := t.Neighbors[N]; ok {
			newTile.Neighbors[S] = v
		}
		if v, ok := t.Neighbors[W]; ok {
			newTile.Neighbors[W] = v
		}
		if v, ok := t.Neighbors[E]; ok {
			newTile.Neighbors[E] = v
		}
	}

	return newTile
}

// overlap ..
func (t *Tile) overlap(edge string, ignoreDirection bool) (direction, bool) {
	invert := invertString(edge)
	for d, v := range t.Edges {
		if v == invert {
			return d, true
		}

		if ignoreDirection {
			if v == edge {
				return d, true
			}
		}
	}

	return 0, false
}

func invertString(s string) string {
	runes := []rune(s)
	out := make([]rune, len(runes))
	for i, v := range runes {
		out[len(runes)-1-i] = v
	}
	return string(out)
}

// ParseInput ..
func ParseInput(f []byte) (map[int]*Tile, error) {
	lf := strings.ReplaceAll(string(f), "\r", "")
	split := strings.Split(lf, "\n\n")

	out := make(map[int]*Tile, len(split))
	mu := &sync.Mutex{}
	eg := &errgroup.Group{}
	re := regexp.MustCompile(`Tile (?P<number>\d+):\n(?P<data>[\.#\n]+)`)
	for _, v := range split {
		eg.Go(func(curr string) func() error {
			return func() error {
				match, err := regutil.MustCaptureNamedGroup(re, []string{"number", "data"}, curr)
				if err != nil {
					return err
				}

				id, err := strconv.Atoi(match["number"])
				if err != nil {
					return err
				}

				dataStr := match["data"]

				dataSplit := strings.Split(dataStr, "\n")

				data := make([][]bool, len(dataSplit))

				for i, v := range dataSplit {
					runes := []rune(v)
					data[i] = make([]bool, len(runes))
					for j, r := range runes {
						data[i][j] = r == '#'
					}
				}

				newTile, err := NewTile(id, data)
				if err != nil {
					return err
				}

				mu.Lock()
				out[id] = newTile
				mu.Unlock()

				return nil
			}
		}(v))
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return out, nil
}

// SolvePart1 ..
func SolvePart1(tiles map[int]*Tile) (int, error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tiles))
	ids := []int{}
	mu := &sync.Mutex{}
	for id, tile := range tiles {
		go func(cID int, cTile *Tile) {
			defer wg.Done()
			count := 0
			for _, theOtherTile := range tiles {
				fits := false
				if cID == theOtherTile.ID {
					continue
				}
				for _, dir := range []direction{N, E, S, W} {
					if _, ok := theOtherTile.overlap(cTile.Edges[dir], true); ok {
						fits = true
						break
					}
				}
				if fits {
					count++
				}
			}
			if count == 2 {
				mu.Lock()
				ids = append(ids, cID)
				mu.Unlock()
			}
		}(id, tile)
	}

	wg.Wait()
	if len(ids) != 4 {
		return 0, fmt.Errorf("expecting 4 corner tiles, got %v", len(ids))
	}

	prod := 1
	for _, v := range ids {
		prod *= v
	}
	return prod, nil
}

// SolvePart2 ..
func SolvePart2(tiles map[int]*Tile) (int, error) {
	// First we assemble the image
	gridSize := int(math.Floor(math.Sqrt(float64(len(tiles)))))
	image := make([][]*Tile, gridSize)
	for i := range image {
		image[i] = make([]*Tile, gridSize)
	}

	// Construct the pool
	wg := &sync.WaitGroup{}
	wg.Add(len(tiles))
	mu := &sync.Mutex{}
	for id, tile := range tiles {
		go func(cID int, cTile *Tile) {
			defer wg.Done()
			neighbours := make(map[direction]int)
			for _, theOtherTile := range tiles {
				if cID == theOtherTile.ID {
					continue
				}
				for _, dir := range []direction{N, E, S, W} {
					if _, ok := theOtherTile.overlap(cTile.Edges[dir], true); ok {
						neighbours[dir] = theOtherTile.ID
					}
				}
			}
			mu.Lock()
			cTile.Neighbors = neighbours
			mu.Unlock()
		}(id, tile)
	}

	wg.Wait()

	// pick TL
	TL := func() *Tile {
		for _, v := range tiles {
			if len(v.Neighbors) == 2 {
				return v
			}
		}
		panic("this should not happen")
	}()

	// Rotate TL
	for !func() bool {
		_, eok := TL.Neighbors[E]
		_, sok := TL.Neighbors[S]
		return eok && sok
	}() {
		*TL = *TL.rotateClockWise()
	}

	image[0][0] = TL

	// Fill up top row
	for i := 1; i < gridSize-1; i++ {
		idNext := image[0][i-1].Neighbors[E]
		pNext := tiles[idNext]
		// rotate until N has no neighbor
		for func() bool {
			_, nok := pNext.Neighbors[N]
			return nok
		}() {
			*pNext = *pNext.rotateClockWise()
		}
		// flip if West was not the last entry
		if pNext.Neighbors[W] != image[0][i-1].ID {
			*pNext = *pNext.mirrorVertical()
		}
		image[0][i] = pNext
	}

	// Find TR
	pTR := tiles[image[0][gridSize-2].Neighbors[E]]
	for !func() bool {
		w, wok := pTR.Neighbors[W]
		return wok && w == image[0][gridSize-2].ID
	}() {
		*pTR = *pTR.rotateClockWise()
	}
	if func() bool {
		_, nok := pTR.Neighbors[N]
		return nok
	}() {
		*pTR = *pTR.mirrorTopDown()
	}

	image[0][gridSize-1] = pTR

	// fill in the rest
	for j := 1; j < gridSize; j++ {
		pL := tiles[image[j-1][0].Neighbors[S]]
		for !func() bool {
			n, nok := pL.Neighbors[N]
			return nok && n == image[j-1][0].ID
		}() {
			*pL = *pL.rotateClockWise()
		}
		if func() bool {
			_, wok := pL.Neighbors[W]
			return wok
		}() {
			*pL = *pL.mirrorVertical()
		}
		image[j][0] = pL
		for i := 1; i < gridSize; i++ {
			pNext := tiles[image[j][i-1].Neighbors[E]]
			for !func() bool {
				n, nok := pNext.Neighbors[N]
				return nok && n == image[j-1][i].ID
			}() {
				*pNext = *pNext.rotateClockWise()
			}
			if w, wok := pNext.Neighbors[W]; !(wok && w == image[j][i-1].ID) {
				*pNext = *pNext.mirrorVertical()
			}
			image[j][i] = pNext
		}
	}

	cellSize := len(TL.Data) - 2

	full := make([][]bool, cellSize*gridSize)
	for i := range full {
		full[i] = make([]bool, cellSize*gridSize)
	}

	for gridv, gridr := range image {
		for rowi, tile := range gridr {
			for i := 0; i < cellSize; i++ {
				for j := 0; j < cellSize; j++ {
					full[cellSize*gridv+i][cellSize*rowi+j] = tile.Data[i+1][j+1]
				}
			}
		}
	}

	fullImage, err := NewTile(0, full)
	if err != nil {
		return 0, err
	}

	seaMonsterStr := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

	seaMonsterSplit := strings.Split(seaMonsterStr, "\n")
	seaMonster := make([][]bool, len(seaMonsterSplit))
	for i, v := range seaMonsterSplit {
		runes := []rune(v)
		seaMonster[i] = make([]bool, len(runes))
		for j, r := range runes {
			seaMonster[i][j] = r == '#'
		}
	}

	smHeight := len(seaMonster)
	smWidth := len(seaMonster[0])

	smCount := 0

	for i := 0; i < 4; i++ {
		fullImage = fullImage.rotateClockWise()
		flip := fullImage.mirrorTopDown()
		for w := 0; w <= cellSize*gridSize-smHeight; w++ {
			for j := 0; j <= cellSize*gridSize-smWidth; j++ {
				isSm := true
				isSmF := true
				for p := 0; p < smHeight; p++ {
					for q := 0; q < smWidth; q++ {
						if isSm && seaMonster[p][q] && !(fullImage.Data[w+p][j+q]) {
							isSm = false
						}
						if isSmF && seaMonster[p][q] && !(flip.Data[w+p][j+q]) {
							isSmF = false
						}
					}
				}
				if isSm {
					smCount++
				}
				if isSmF {
					smCount++
				}
			}
		}
	}

	smSize := 0

	for _, r := range seaMonster {
		for _, v := range r {
			if v {
				smSize++
			}
		}
	}

	totalHash := 0
	for _, r := range fullImage.Data {
		for _, v := range r {
			if v {
				totalHash++
			}
		}
	}

	return totalHash - smSize*smCount, nil
}
