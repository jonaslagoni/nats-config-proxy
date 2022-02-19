package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats-config-proxy/internal/server"
)

const usageStr = `
Options:
    -d, --dir <directory>         Directory for storing data (default is the current directory.)
    -s, --snapshot <name>         Take snapshot of the configuration
    -h, --help                    Show this message
    -v, --version                 Show version
`

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: nats-rest-config-snapshot [options...]")
		fmt.Fprintln(os.Stderr, usageStr)
	}

	var (
		showVersion  bool
		showHelp     bool
		snapshotName string
	)

	opts := &server.Options{
		NoLog: true,
	}
	flag.StringVar(&opts.DataDir, "data", ".", "Directory for storing data.")
	flag.StringVar(&opts.DataDir, "dir", ".", "Directory for storing data.")
	flag.StringVar(&opts.DataDir, "d", ".", "Directory for storing data.")
	flag.StringVar(&snapshotName, "s", server.DefaultSnapshotName, "Snapshot of the configuration.")
	flag.StringVar(&snapshotName, "snapshot", server.DefaultSnapshotName, "Snapshot of the configuration.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.BoolVar(&showVersion, "version", false, "Show version.")
	flag.BoolVar(&showHelp, "h", false, "Show help.")
	flag.BoolVar(&showHelp, "help", false, "Show help.")
	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		os.Exit(0)
	case showVersion:
		fmt.Printf("nats-rest-config-snapshot %s\n", server.Version)
		os.Exit(0)
	}

	s := server.NewHttpServer(opts)
	fmt.Printf("Taking %q snapshot...\n", snapshotName)
	if err := s.TakeSnapshot(snapshotName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "OK")
}
