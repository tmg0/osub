package console

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"osub/pkg/parser"
	"osub/pkg/resolve"
	"osub/pkg/shared"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the osub service",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := shared.ReadConfig()

		if err != nil {
			fmt.Println("Cannot read osub config json file: ", err)
		}

		for {
			for _, sub := range conf.Subscriptions {
				resp, err := http.Get(sub.URL)

				if err != nil {
					fmt.Println("Error fetching subscription: ", err)
				}

				body, err := io.ReadAll(resp.Body)

				if err != nil {
					log.Fatalf("Error request subscription link: %v", err)
				}

				links, err := parser.Subscription(string(body))

				if err != nil {
					log.Fatalf("Error parsing Subscription link: %v", err)
				}

				for _, link := range links {
					if strings.HasPrefix(link, shared.VMESS_PREFIX) {
						config, err := parser.Vmess(link)

						if err != nil {
							log.Fatalf("Error parsing Vmess link: %v", err)
						}

						fmt.Println(config)
					}
				}

				err = resp.Body.Close()

				if err != nil {
					log.Fatalf("Response closed error: %v", err)
				}
			}

			duration, err := resolve.Interval(conf.Interval)

			if err != nil {
				log.Fatalf("Error parsing Interval string: %v", err)
			}

			time.Sleep(duration)
		}
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
