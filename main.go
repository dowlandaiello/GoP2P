package main

import (
	"flag"
	"net/http"

	"github.com/mitsukomegumi/GoP2P/cli"
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/rpc/environment"
	handler "github.com/mitsukomegumi/GoP2P/rpc/handler"
	node "github.com/mitsukomegumi/GoP2P/rpc/node"
	environmentProto "github.com/mitsukomegumi/GoP2P/rpc/proto/environment"
	handlerProto "github.com/mitsukomegumi/GoP2P/rpc/proto/handler"
	nodeProto "github.com/mitsukomegumi/GoP2P/rpc/proto/node"
	upnpProto "github.com/mitsukomegumi/GoP2P/rpc/proto/upnp"
	upnpServer "github.com/mitsukomegumi/GoP2P/rpc/upnp"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

var (
	terminalFlag = flag.Bool("terminal", false, "launch GoP2P in terminal mode")                   // Init term flag
	upnpFlag     = flag.Bool("upnp", false, "launch GoP2P without automatic UPnP port forwarding") // Init upnp flag
)

func main() {
	flag.Parse() // Parse flags

	if !*upnpFlag {
		go upnp.ForwardPortSilent(8080) // Forward port 8080
	}

	startRPCServer() // Start RPC server

	if *terminalFlag {
		cli.NewTerminal() // Initialize terminal
	}

	startNode() // Attempt to start GoP2P in node mode

	go common.Forever() // Prevent main from closing
	select {}           // Prevent main from closing
}

// startRPCServer - start RPC server
func startRPCServer() {
	nodeHandler := nodeProto.NewNodeServer(&node.Server{}, nil)                             // Init handler
	handlerHandler := handlerProto.NewHandlerServer(&handler.Server{}, nil)                 // Init handler
	environmentHandler := environmentProto.NewEnvironmentServer(&environment.Server{}, nil) // Init handler
	upnpHandler := upnpProto.NewUpnpServer(&upnpServer.Server{}, nil)                       // Init handler

	mux := http.NewServeMux() // Init mux

	mux.Handle(nodeProto.NodePathPrefix, nodeHandler)                      // Start mux node handler
	mux.Handle(handlerProto.HandlerPathPrefix, handlerHandler)             // Start mux handler handler
	mux.Handle(environmentProto.EnvironmentPathPrefix, environmentHandler) // Start mux environment handler
	mux.Handle(upnpProto.UpnpPathPrefix, upnpHandler)                      // Start mux upnp handler

	go http.ListenAndServe(":8080", mux) // Start server
}

// startNode - attempt to execute attachnode, starthandler commands
func startNode() {
	terminal := cli.Terminal{Variables: []cli.Variable{}} // Init terminal

	terminal.HandleCommand("node.Attach()") // Attach node

	terminal.HandleCommand("node.StartHandler()") // Start handler
}

/* TODO:
- Add custom RPC addresses
- On attachment, make sure to write changes to a discovered node's db.
- Add receiver usage in RPC client
*/
