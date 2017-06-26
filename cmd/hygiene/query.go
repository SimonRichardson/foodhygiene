package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/SimonRichardson/foodhygiene/pkg/query"
	"github.com/SimonRichardson/foodhygiene/pkg/service"
	"github.com/SimonRichardson/foodhygiene/pkg/ui"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const (
	// APIRatingsFoodVersion states which api version should be used for
	// http://api.ratings.food.gov.uk
	APIRatingsFoodVersion = 2

	// APIRatingsFoodURL states which url should be used for accessing the API
	APIRatingsFoodURL = "http://api.ratings.food.gov.uk"
)

const (
	defaultCache = true
)

// runQuery creates all the dependencies required to create and run the query
// end point for the cipher component.
func runQuery(args []string) error {
	var (
		flagset = flag.NewFlagSet("query", flag.ExitOnError)

		debug   = flagset.Bool("debug", false, "debug logging")
		apiAddr = flagset.String("api", defaultAPIAddr, "listen address for ingest and store APIs")
		cache   = flagset.Bool("cache", defaultCache, "use cached results for better responsiveness")
		uiLocal = flagset.Bool("ui.local", false, "Ignores embedded files and goes straight to the filesystem")
	)

	flagset.Usage = usageFor(flagset, "query [flags]")
	if err := flagset.Parse(args); err != nil {
		return nil
	}

	// Setup the logger.
	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		if *debug {
			logLevel = level.AllowAll()
		}
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = level.NewFilter(logger, logLevel)
	}

	// Parse the apiNetwork and apiAddress from the flag set
	apiNetwork, apiAddress, err := parseAddr(*apiAddr, defaultAPIPort)
	if err != nil {
		return err
	}

	// Create the api listener for the service
	apiListener, err := net.Listen(apiNetwork, apiAddress)
	if err != nil {
		return err
	}
	level.Debug(logger).Log("API", fmt.Sprintf("%s://%s", apiNetwork, apiAddress))

	// Execution group.
	defer apiListener.Close()

	// Service wraps the food agency API
	serv := service.New(APIRatingsFoodURL, APIRatingsFoodVersion, log.With(logger, "component", "service"))
	if *cache {
		serv = service.NewCache(serv)
	}

	// API that is going to handle the incoming requests.
	api := query.NewAPI(serv, log.With(logger, "component", "api"))

	mux := http.NewServeMux()
	mux.Handle("/query/", http.StripPrefix("/query", api))
	mux.Handle("/ui/", ui.NewAPI(*uiLocal, log.With(logger, "component", "ui")))

	return http.Serve(apiListener, mux)

}
