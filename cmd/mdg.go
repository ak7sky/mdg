/*
Copyright © 2023 ak7sky
*/

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ak7sky/mdg/gen"
	"github.com/spf13/cobra"
)

var (
	nOfS int
	nOfT int
	file string
)

var mdgCmd = &cobra.Command{
	Use:   "mdg",
	Short: "Tool to generate MD template",
	Run: func(cmd *cobra.Command, args []string) {
		wr, err := gen.GetWriterForMd(file)
		if err != nil {
			log.Fatal(err)
		}
		err = gen.BuildMd(nOfS, nOfT, wr)
		if err != nil {
			log.Fatal(err)
		}
		if file != "" {
			fmt.Printf("%s.md was created\n", file)
		}
	},
}

func Execute() {
	err := mdgCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	mdgCmd.Flags().IntVarP(&nOfS, "nofs", "s", 1, "number of sections in MD template")
	mdgCmd.Flags().IntVarP(&nOfT, "noft", "t", 1, "number of topics in each section")
	mdgCmd.Flags().StringVarP(&file, "tofile", "f", "", "name of file to write MD template")
}
