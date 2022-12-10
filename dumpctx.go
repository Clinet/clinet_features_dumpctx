package dumpctx

import (
	"github.com/Clinet/clinet_cmds"
	"github.com/Clinet/clinet_features"
)

var Feature = features.Feature{
	Name: "dumpctx",
	Desc: "Debug command provided by Clinet.",
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