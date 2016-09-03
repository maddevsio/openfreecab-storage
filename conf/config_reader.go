package conf

import (
	"os"

	"github.com/gen1us2k/log"
	"github.com/urfave/cli"
)

// Version stores current service version
var (
	Version      string
	HTTPBindAddr string
	BaseURL      string
	LogLevel     string
	TestMode     bool
)

type Configuration struct {
	data *StorageConfig
	app  *cli.App
}

// NewConfigurator is constructor and creates a new copy of Configuration
func NewConfigurator() *Configuration {
	Version = "0.1dev"
	app := cli.NewApp()
	app.Name = "Open free cabs storage"
	app.Usage = "Storage for drivers"
	return &Configuration{
		data: &StorageConfig{},
		app:  app,
	}
}

func (c *Configuration) fillConfig() *StorageConfig {
	return &StorageConfig{
		HTTPBindAddr: HTTPBindAddr,
		BaseURL:      BaseURL,
		TestMode:     TestMode,
	}
}
func (c *Configuration) After(cb func(c *cli.Context) error) {
	c.app.After = func(ctx *cli.Context) error {
		log.SetLevel(log.MustParseLevel(LogLevel))
		c.data = c.fillConfig()
		return cb(ctx)
	}
}

// Run is wrapper around cli.App
func (c *Configuration) Run() error {
	c.app.Before = func(ctx *cli.Context) error {
		log.SetLevel(log.MustParseLevel(LogLevel))
		return nil
	}
	c.app.Flags = c.setupFlags()
	return c.app.Run(os.Args)
}

// App is public method for Configuration.app
func (c *Configuration) App() *cli.App {
	return c.app
}

func (c *Configuration) setupFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "http_bind_addr",
			Value:       ":8090",
			Usage:       "Define custom http port to bind to",
			EnvVar:      "HTTP_BIND_ADDR",
			Destination: &HTTPBindAddr,
		},
		cli.StringFlag{
			Name:        "base_url",
			Value:       "http://localhost:8090",
			Usage:       "Define custom base url for project",
			EnvVar:      "BASE_URL",
			Destination: &BaseURL,
		},
		cli.StringFlag{
			Name:        "loglevel",
			Value:       "debug",
			Usage:       "set log level",
			Destination: &LogLevel,
			EnvVar:      "LOG_LEVEL",
		},
		cli.BoolFlag{
			Name:        "test_mode",
			Usage:       "set test mode",
			Destination: &TestMode,
			EnvVar:      "TEST_MODE",
		},
	}

}

// Get returns filled StorageConfig
func (c *Configuration) Get() *StorageConfig {
	c.data = c.fillConfig()
	return c.data
}
