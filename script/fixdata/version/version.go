package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/6 18:22
 * @Desc:
 */

const (
	version = "v0.1.0"
)

func New() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Long:  `Print the version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("fixdata version %s \n", version)
		},
	}
	return versionCmd
}
