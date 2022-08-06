package experimental

import (
	"fmt"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/filter"
	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/mesh"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	availableSubcommands []*cobra.Command
)

// ExpCmd represents the Performance Management CLI command
var ExpCmd = &cobra.Command{
	Use:   "exp",
	Short: "Experimental commands for mesheryctl",
	Long:  `List of experimental commands for testing and evaluation purpose.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// return errors.New(utils.ExpError(fmt.Sprintln("mesheryctl exp requires at least 1 arg(s), got only 0")))
			return cmd.Help()

		}
		if ok := utils.IsValidSubcommand(availableSubcommands, args[0]); !ok {
			return errors.New(utils.ExpError(fmt.Sprintf("'%s' is a invalid command. Use 'mesheryctl exp --help' to display usage guide.'\n", args[0])))
		}
		return nil
	},
}

func init() {
	availableSubcommands = []*cobra.Command{mesh.MeshCmd, filter.FilterCmd}
	ExpCmd.AddCommand(availableSubcommands...)
}
