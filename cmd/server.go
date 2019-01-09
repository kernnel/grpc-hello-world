package cmd

import (
	"github.com/spf13/cobra"
	"grpc-hello-world/server/hello"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50056", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "conf/server/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "conf/server/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-name", "", "grpc-hello-world", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto/hello", "path to the directory which contains swagger definitions")

	rootCmd.AddCommand(serverCmd)
}
