package console

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
  Short: "Run the osub service",
  Run: func(cmd *cobra.Command, args []string) {
		conf, err := shared.ReadConfig()

		if err != nil {
			fmt.Println("Cannot read osub config json file: ", err)
		}

		for _, sub := range conf.Subscriptions {
			resp, err := http.Get(sub.URL)

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
		}
  },
}

func init() {
  RootCmd.AddCommand(RunCmd)
}
