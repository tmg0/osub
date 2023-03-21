package cmd

import (
	"fmt"
	"io"
	"log"
  "strings"
	"net/http"
  "github.com/spf13/cobra"
	"osub/pkg/parser"
	"osub/pkg/shared"
)

var RunCmd = &cobra.Command{
  Use:   "run",
  Short: "Runs the app",
  Long:  `This command runs the app and performs some awesome things. For example:`,
  Run: func(cmd *cobra.Command, args []string) {
    resp, err := http.Get("")

		if err != nil {
			fmt.Println("Error fetching subscription: ", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalf("Error request subscription link: %v", err)
		}

		links, err := parser.Subscription(string(body))

		if err != nil {
			log.Fatalf("Error parsing Subscription link: %v", err)
		}

		for _, link := range *links {
			if strings.HasPrefix(link, shared.VMESS_PREFIX) {
				config, err := parser.Vmess(link)

				if err != nil {
					log.Fatalf("Error parsing Vmess link: %v", err)
				}

				fmt.Println(config)
			}
		}
  },
}

func init() {
  RootCmd.AddCommand(RunCmd)
}
