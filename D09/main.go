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

func compressFSV1(fs []int) {
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
			break
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

func compressViaDiskMap(dm []DiskItem) {
	for ix := len(dm) - 1; ix >= 0; ix-- {
		if dm[ix].id == FS_SPACE {
			continue
		}

		i := slices.IndexFunc(dm, func(it DiskItem) bool {
			return it.id == FS_SPACE && it.size >= dm[ix].size
		})

		// No space found
		if i == -1 {
			continue
		}

		// TODO: check above again...dont think this is ok yet
	}
}

func main() {

	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	dm := parseDiskMap(f)

	fs := makeFilesystem(dm)

	fs1 := slices.Clone(fs)
	compressFSV1(fs1)
	fmt.Println("Part1: ", calcFSChecksum(fs1))

	dm2 := slices.Clone(dm)
	compressViaDiskMap(dm2)
	fmt.Println("Part2: ", calcFSChecksum(makeFilesystem(dm2)))
}
