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

func parseFileSysAndDiskMap(f *os.File) (fs, dm []int) {
	sc := bufio.NewScanner(f)

	diskMap := ""

	if ok := sc.Scan(); !ok {
		panic("No line read")
	} else {
		diskMap = sc.Text()
	}

	isFile := true

	fs, dm = []int{}, []int{}

	id := 0
	for _, numstr := range diskMap {
		sizeFsItem, _ := strconv.ParseInt(string(numstr), 10, 64)

		dm = append(dm, int(sizeFsItem))

		descriptor := 0

		if isFile {
			descriptor = id
			id = id + 1
		} else {
			descriptor = FS_SPACE
		}

		isFile = !isFile

		for range iter.Interval(0, sizeFsItem) {
			fs = append(fs, descriptor)
		}
	}

	return fs, dm
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
			break
		}

		cs = cs + ix*id
	}

	return cs
}

func main() {

	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fs, _ := parseFileSysAndDiskMap(f)
	// fmt.Println(dm)

	compressFS(fs)
	fmt.Println(calcFSChecksum(fs))
}
