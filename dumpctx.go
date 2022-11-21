package dumpctx

import (
	"github.com/Clinet/clinet_cmds"
	"github.com/Clinet/clinet_features"
)

var Feature = features.Feature{
	Help: "Use `/dumpctx sub1` or `/dumpctx sub2` to dump Clinet's subcommand context.",
	Name: "dumpctx",
	Cmds: []*cmds.Cmd{
		cmds.NewCmd("dumpctx", "DEBUG: Dumps command context as built by Clinet", nil).AddSubCmds(
			cmds.NewCmd("sub1", "Test subcommand 1", handleDumpCtx),
			cmds.NewCmd("sub2", "Test subcommand 2", handleDumpCtx),
		),
	},
}

func handleDumpCtx(ctx *cmds.CmdCtx) *cmds.CmdResp {
	return cmds.NewCmdRespEmbed("Dump of ctx (*cmds.CmdCtx)", "```JSON\n" + ctx.String() + "```")
}