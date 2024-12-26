package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type RootCmdData struct {
	fileLeft  string
	fileRight string
}

func (s *RootCmdData) runCmd() error {
	lfile, lerr := os.Open(s.fileLeft)

	if lerr != nil {
		return lerr
	}

	defer lfile.Close()

	rfile, rerr := os.Open(s.fileRight)

	if rerr != nil {
		return rerr
	}

	defer rfile.Close()

	lscanner, rscanner := bufio.NewScanner(lfile), bufio.NewScanner(rfile)

	llist, rlist := make([]uint, 0, 10), make([]uint, 0, 10)

	for lscanner.Scan() {
		v, err := strconv.ParseUint(lscanner.Text(), 10, 64)

		if err != nil {
			return err
		}

		llist = append(llist, uint(v))
	}

	for rscanner.Scan() {
		v, err := strconv.ParseUint(rscanner.Text(), 10, 64)

		if err != nil {
			return err
		}

		rlist = append(rlist, uint(v))
	}

	if len(llist) != len(rlist) {
		return fmt.Errorf("length of left list and right list do not match (left = %d, right = %d items)", len(llist), len(rlist))
	}

	fmt.Printf("Total distance between both lists is %d", GetTotalDistance(llist, rlist))

	return nil
}

var rootCmdData = RootCmdData{fileLeft: "", fileRight: ""}

var rootCmd = &cobra.Command{
	Use:   "d01 [leftList] [rightList]",
	Short: "CLI for Advent of Code 2024 - Day 01\n",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		rootCmdData.fileLeft = args[0]
		rootCmdData.fileRight = args[1]
		if err := rootCmdData.runCmd(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func Run() {
	if e := rootCmd.Execute(); e != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: '%s'\n", e)
		os.Exit(1)
	}
}
