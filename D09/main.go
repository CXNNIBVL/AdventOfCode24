package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/CXNNIBVL/goutil/iter"
)

const (
	FILE      = "../inputs/D09/input"
	TEST_FILE = "../inputs/D09/input_test"
)

const (
	FS_SPACE = -1
)

func makeFilesystem(dm []DiskItem) (fs []int) {
	fs = []int{}

	for _, fsitem := range dm {
		for range iter.Interval(0, fsitem.size) {
			fs = append(fs, fsitem.id)
		}
	}

	return fs
}

func moveToNextSpace(fs []int) []int {
	ix := slices.Index(fs, FS_SPACE)
	return fs[ix:]
}

func moveToNextChunk(fs []int) ([]int, bool) {
	for ix := len(fs) - 1; ix >= 0; ix-- {
		if fs[ix] != FS_SPACE {
			return fs[:ix+1], true
		}
	}

	return fs, false
}

func compressFS(fs []int) {
	for len(fs) != 1 {
		fs = moveToNextSpace(fs)
		_fs, hasNext := moveToNextChunk(fs)
		fs = _fs

		if !hasNext {
			break
		}

		fs[0] = fs[len(fs)-1]
		fs[len(fs)-1] = FS_SPACE
	}
}

func calcFSChecksum(fs []int) int {
	cs := 0

	for ix, id := range fs {
		if id == FS_SPACE {
			continue
		}

		cs = cs + ix*id
	}

	return cs
}

type DiskItem struct {
	id, size int
}

func parseDiskMap(f *os.File) []DiskItem {
	sc := bufio.NewScanner(f)

	diskMap := ""

	if ok := sc.Scan(); !ok {
		panic("No line read")
	} else {
		diskMap = sc.Text()
	}

	isFile := true

	dm := []DiskItem{}

	id := 0
	for _, numstr := range diskMap {
		sizeFsItem, _ := strconv.ParseInt(string(numstr), 10, 64)

		descriptor := 0

		if isFile {
			descriptor = id
			id = id + 1
		} else {
			descriptor = FS_SPACE
		}

		dm = append(dm, DiskItem{id: descriptor, size: int(sizeFsItem)})
		isFile = !isFile
	}

	return dm
}

func compressViaDiskMap(dm []DiskItem) []DiskItem {
	ix := len(dm) - 1
	for ix >= 0 {

		if dm[ix].id == FS_SPACE {
			ix--
			continue
		}

		// Find space that can fit our block
		i := slices.IndexFunc(dm[:ix], func(it DiskItem) bool {
			return it.id == FS_SPACE && it.size >= dm[ix].size
		})

		// No space found or not moveable to the left
		if i == -1 {
			ix--
			continue
		}

		sizediff := dm[i].size - dm[ix].size

		// Swap
		dm[i] = dm[ix]
		dm[ix].id = FS_SPACE

		if sizediff > 0 {
			// Merge spaces
			if it := &dm[i+1]; it.id == FS_SPACE {
				it.size = it.size + sizediff
			} else {
				// Append space block
				dm = slices.Insert(dm, i+1, DiskItem{id: FS_SPACE, size: sizediff})
				ix = ix + 1
			}
		}

		ix--
	}

	return dm
}

func main() {

	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	dm := parseDiskMap(f)

	dm1 := slices.Clone(dm)
	fs1 := makeFilesystem(dm1)
	compressFS(fs1)
	fmt.Println("Part1: ", calcFSChecksum(fs1))

	dm2 := slices.Clone(dm)
	fmt.Println("Part2: ", calcFSChecksum(makeFilesystem(compressViaDiskMap(dm2))))
}
