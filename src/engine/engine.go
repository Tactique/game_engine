package engine

import (
	"flag"
	"github.com/Tactique/golib/logger"
	"os"
)

func Main() {
	var port = flag.Int("port", 5269, "path to write the generated json")
	var write_template = flag.Bool("write_templates", false, "whether to write the template response json")
	var template_out_dir = flag.String("template_out_dir", ".", "path to write the generated template json")
	var logpath = flag.String("logpath", "/dev/stdout", "Logging location")
	flag.Parse()

	logger.SetupLoggerHelper(*logpath)

	if *write_template {
		generateTemplates(*template_out_dir)
		os.Exit(0)
	} else {
		logger.Infof("Starting Tactique on port %d!", *port)
		panic(listenForever(*port))
	}
}
