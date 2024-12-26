package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type GenCmdData struct {
	idrange   *uint
	itemcount *uint
	pathLeft  string
	pathRight string
}

var genCmdData = GenCmdData{idrange: nil, itemcount: nil}

func (s *GenCmdData) runCmd() error {
	lfile, lerr := os.Create(s.pathLeft)

	if lerr != nil {
		return lerr
	}

	defer lfile.Close()

	rfile, rerr := os.Create(s.pathRight)

	if rerr != nil {
		return rerr
	}

	defer rfile.Close()

	llist, rlist := makeList(*s.itemcount, *s.idrange), makeList(*s.itemcount, *s.idrange)

	for _, v := range llist[0 : len(llist)-1] {
		fmt.Fprintln(lfile, v)
	}
	fmt.Fprint(lfile, llist[len(llist)-1])

	for _, v := range rlist[0 : len(rlist)-1] {
		fmt.Fprintln(rfile, v)
	}
	fmt.Fprint(rfile, rlist[len(rlist)-1])

	return nil
}

var genCmd = &cobra.Command{
	Use:   "gen [fileLeft] [fileRight]",
	Short: "Generates the left and right lists. Overwrites them if they already exist.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		genCmdData.pathLeft = args[0]
		genCmdData.pathRight = args[1]

		if err := genCmdData.runCmd(); err != nil {
			fmt.Fprintf(os.Stderr, "error occurred: %s", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	genCmdData.idrange = genCmd.Flags().UintP("idrange", "i", 30, "Defines the ID bound to be between 0...VALUE")
	genCmdData.itemcount = genCmd.Flags().UintP("count", "c", 10, "Defines the item count generated in the list")
	rootCmd.AddCommand(genCmd)
}
