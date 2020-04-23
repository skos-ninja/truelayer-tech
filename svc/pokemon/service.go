package pokemon

import (
	"fmt"
	"net/http"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/app"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/rpc"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:  "pokemon",
	RunE: runE,
}

func runE(cmd *cobra.Command, args []string) error {
	// We use https://github.com/gin-gonic/gin here as this is
	// the router I have typically used and the overhead compared
	// to using net/http directly is negligible compared to the cost
	// of development of the features gin provides.
	//
	// The default gin engine includes request logging out of the box
	// and also includes a standard recovery handler for panics.
	//
	// In previous implementations of gin by me it has been abstracted
	// away to provide standard functionality across all microservices.
	// This would include standard auth handling, standardised error responses,
	// error reporting to places like sentry.io and metric tracking using prometheus.
	//
	// Due to the nature of this service being just an example I have not included
	// this functionality as it would increase the time of development here whilst
	// also not being required by the tech test.
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	app := app.New()
	rpc := rpc.New(app)

	r.GET("/pokemon/:id", rpc.GetPokemon)

	// As this is designed to be run inside a container we should only
	// allow binding to 0.0.0.0 due to how networking is done within docker.
	//
	// A good article explaining this can be found here:
	// https://pythonspeed.com/articles/docker-connection-refused/
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}
	return r.Run(fmt.Sprintf("0.0.0.0:%v", port))
}
