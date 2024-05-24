package main

import (
	"os"
	"path/filepath"

	nm "github.com/scalarorg/abci/node"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"
	"github.com/cometbft/cometbft/cmd/cometbft/commands/debug"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/cli"
)

func main() {
	rootCmd := RootCmd
	rootCmd.AddCommand(
		cmtcmd.GenValidatorCmd,
		cmtcmd.InitFilesCmd,
		cmtcmd.LightCmd,
		cmtcmd.ReplayCmd,
		cmtcmd.ReplayConsoleCmd,
		cmtcmd.ResetAllCmd,
		cmtcmd.ResetPrivValidatorCmd,
		cmtcmd.ResetStateCmd,
		cmtcmd.ShowValidatorCmd,
		cmtcmd.TestnetFilesCmd,
		cmtcmd.ShowNodeIDCmd,
		cmtcmd.GenNodeKeyCmd,
		cmtcmd.VersionCmd,
		cmtcmd.RollbackStateCmd,
		cmtcmd.CompactGoLevelDBCmd,
		debug.DebugCmd,
		cli.NewCompletionCmd(rootCmd, true),
	)

	// NOTE:
	// Users wishing to:
	//	* Use an external signer for their validators
	//	* Supply an in-proc abci app
	//	* Supply a genesis doc file from another source
	//	* Provide their own DB implementation
	// can copy this file and use something other than the
	// DefaultNewNode function
	nodeFunc := nm.DefaultNewNode

	// Create & start node
	rootCmd.AddCommand(nm.NewRunNodeCmd(nodeFunc))

	cmd := cli.PrepareBaseCmd(rootCmd, "CMT", os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultTendermintDir)))
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
