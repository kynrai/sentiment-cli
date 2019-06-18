package sweep

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kynrai/sentiment-cli/pkg/twitter"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sweep",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sweep cmd")
		r := twitter.New()
		resp, err := r.Tweets30Days(context.Background(), "", 100, time.Now().AddDate(0, -1, 0), time.Now())
		if err != nil {
			log.Fatal(err)
		}
		for _, res := range resp.Results {
			fmt.Println(res.Text)
		}
	},
}
