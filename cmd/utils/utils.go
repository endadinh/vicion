package utils

import (
	"github.com/vicion/vicion/eth"
	"github.com/vicion/vicion/eth/downloader"
	"github.com/vicion/vicion/ethstats"
	"github.com/vicion/vicion/les"
	"github.com/vicion/vicion/node"
	"github.com/vicion/vicion/vicionx"
	"github.com/vicion/vicion/vicionxlending"
	whisper "github.com/vicion/vicion/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var vicionXServ *vicionx.VicionX
			ctx.Service(&vicionXServ)
			var lendingServ *vicionxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, vicionXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterVicionXService(stack *node.Node, cfg *vicionx.Config) {
	vicionX := vicionx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return vicionX, nil
	}); err != nil {
		Fatalf("Failed to register the VicionX service: %v", err)
	}

	// register vicionxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return vicionxlending.New(vicionX), nil
	}); err != nil {
		Fatalf("Failed to register the VicionXLending service: %v", err)
	}
}
