package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/weiguoyu/route/read"
	"github.com/weiguoyu/route/exec"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "ENV_DEBUG",
		Name:   "debug",
		Usage:  "enable debug mode",
	},
	cli.StringFlag{
		EnvVar: "IP_ADDRESS_FILEPATH",
		Name:   "ip_address_filepath",
		Value:  "./IPAddressTest.ini",
		Usage:  "IP Address filepath",
	},
	cli.StringFlag{
		EnvVar: "GATEWAY_IP",
		Name:   "gateway_ip",
		Value:  "10.88.88.2",
		Usage:  "gateway ip",
	},
	cli.StringFlag{
		EnvVar: "INTERFACE",
		Name:   "interface",
		Value:  "IF 5",
		Usage:  "the gateway device interface",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "route"
	app.Usage = "add route for windows"
	app.Flags = flags
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		logrus.Error("app.Run err: ", err)
	}
}

func run(c *cli.Context) error {
	// set logger
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
	}
	logrus.Info("app start ...")
	addRoute(c)
	
	return nil
}

func addRoute(c *cli.Context) error {
	// add route
	ip_map, err := read.ReadFile(c.String("ip_address_filepath"))
	if err != nil {
		logrus.Errorf("err: %v", err)
		return err
	}
	for ip, mask := range ip_map {
		exec.ExecCmd("route", "add", ip, "MASK", mask.(string), c.String("gateway_ip"))
	}
	return err
}
