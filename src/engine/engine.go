package engine

import (
	"flag"
	"github.com/Tactique/golib/logger"
)

func Main() {
	var port = flag.Int("port", 5269, "path to write the generated json")
	flag.Parse()

	logger.SetupLoggerHelper("engine.log")

	logger.Infof("Starting Tactique on port %d!", *port)
	ListenForever(*port)
}
