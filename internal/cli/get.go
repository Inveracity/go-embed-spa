package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/inveracity/go-embed-spa/internal/server"
)

type Get struct {
}

func (g *Get) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get data from server",

		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := g.getData()
			fmt.Println(result)
			return err
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}

func (g *Get) getData() (string, error) {
	resp, err := http.Get("http://localhost:3000/api")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var hello server.Hello
	if err := json.Unmarshal(body, &hello); err != nil {
		return "", err
	}
	ret := fmt.Sprintf("Hello %s", hello.Hello)
	return ret, nil
}
