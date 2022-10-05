package cmd

import (
	"tinyurl/pkg/api"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// ApiServerCmd api server init point
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE:  RunServerCmd,
}

func RunServerCmd(cmd *cobra.Command, args []string) error {
	app := fiber.New()
	api.SetRoutes(app)
	app.Listen(":3000")

	return nil
}
