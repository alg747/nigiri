package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var lnd = cli.Command{
	Name:   "lnd",
	Usage:  "invoke LND command line",
	Action: lndAction,
}

func lndAction(ctx *cli.Context) error {

	if isRunning, _ := nigiriState.GetBool("running"); !isRunning {
		return errors.New("nigiri is not running")
	}

	network, err := nigiriState.GetString("network")
	if err != nil {
		return err
	}

	rpcArgs := []string{"exec", "lnd", "lncli", "--network=" + network}
	cmdArgs := append(rpcArgs, ctx.Args().Slice()...)
	bashCmd := exec.Command("docker", cmdArgs...)
	bashCmd.Stdout = os.Stdout
	bashCmd.Stderr = os.Stderr

	if err := bashCmd.Run(); err != nil {
		return err
	}

	return nil
}