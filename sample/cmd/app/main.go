package main

import (
	"fmt"

	"github.com/kgmedia-data/gaia/pkg/handler"
	"github.com/kgmedia-data/gaia/sample/config"
	"github.com/kgmedia-data/gaia/sample/internal/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configPath string

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rootCmd := &cobra.Command{
		Use:   "main",
		Short: "A sample gaia application",
		Long:  "A sample gaia application built with Go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to gaia!")
		},
	}

	rootCmd.PersistentFlags().StringVar(&configPath, "config",
		"config/config.yaml", "path for the config")

	rootCmd.AddCommand(appCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func appCmd() *cobra.Command {
	var (
		rest bool
	)

	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Run the gaia application",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig[config.Config](configPath)
			if err != nil {
				logrus.Fatalln(err)
			}
			app := app.NewApp(cfg)

			registry := handler.NewRegistry()
			registry.Register("METRIC", handler.NewMetricServer(cfg.Metric.Host))

			if rest {
				registry.Register("REST", app.CreateRestServer())
			}

			registry.StartAll()
			registry.StopAll()
		},
	}

	appCmd.Flags().BoolVar(&rest, "rest", false, "enable REST API")

	return appCmd
}
