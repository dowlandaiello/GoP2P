package main

import (
	"flag"

	"github.com/mitsukomegumi/GoP2P/cli"
	"github.com/mitsukomegumi/GoP2P/common"
)

var (
	terminalFlag = flag.Bool("terminal", false, "launch gop2p in terminal mode") // Init term flag
)

func main() {
	flag.Parse() // Parse flags

	if *terminalFlag {
		cli.NewTerminal() // TODO: on attachment, make sure to write changes to a discovered node's db.
	}

	startNode() // Attempt to start gop2p in node mode

	go common.Forever() // Prevent main from closing
	select {}           // Prevent main from closing
}

// startNode - attempt to execute attachnode, starthandler commands
func startNode() {
	terminal := cli.Terminal{Variables: []cli.Variable{}} // Init terminal

	terminal.HandleCommand("node.Attach()") // Attach node

	terminal.HandleCommand("node.StartHandler()") // Start handler
}
