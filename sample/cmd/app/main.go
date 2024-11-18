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
		rest           bool
		employeeJob    bool
		employeeStream bool
		useMem         bool
		gcpLogging     bool
	)

	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Run the gaia application",
		Run: func(cmd *cobra.Command, args []string) {

			registry := handler.NewRegistry()

			if useMem {
				app := app.NewAppMem()
				if rest {
					registry.Register("REST", app.CreateRestServer())
				}
				if employeeStream {
					registry.Register("EMPLOYEE-STREAM", app.CreateEmployeeStream())
				}
				if employeeJob {
					job := app.CreateEmployeeJob()
					job.Run()
				}
			} else {
				cfg, err := config.LoadConfig[config.Config](configPath)
				if err != nil {
					logrus.Fatalln(err)
				}
				app := app.NewApp(cfg)

				registry.Register("METRIC", handler.NewMetricServer(cfg.Metric.Host))
				if rest {
					registry.Register("REST", app.CreateRestServer())
				}

				if gcpLogging {
					registry.Register("GCP LOGGING", app.CreateLoggerStream())
				}
			}

			registry.StartAll()
			registry.StopAll()
		},
	}

	appCmd.Flags().BoolVar(&rest, "rest", false, "enable REST API")
	appCmd.Flags().BoolVar(&employeeJob, "employee-job", false, "enable employee job")
	appCmd.Flags().BoolVar(&employeeStream, "employee-stream", false, "enable employee stream")
	appCmd.Flags().BoolVar(&useMem, "use-mem", false, "use memory repository")
	appCmd.Flags().BoolVar(&gcpLogging, "gcp-log", false, "use GCP Logging")

	return appCmd
}
