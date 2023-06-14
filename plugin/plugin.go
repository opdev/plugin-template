package plugin

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/go-logr/logr"
	"github.com/opdev/knex/plugin/v0"
	"github.com/opdev/knex/types"
	"github.com/opdev/plugin-template/internal/flags"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Assert that we implement the Plugin interface.
var _ plugin.Plugin = NewPlugin()

var vers = semver.MustParse("0.0.1")

func init() {
	plugin.Register("example-plugin", NewPlugin())
}

type plug struct{}

func NewPlugin() *plug {
	p := plug{}
	// plugin-related things may happen here.
	return &p
}

func (p *plug) Register() error {
	return nil
}

func (p *plug) Name() string {
	return "Example Plugin"
}

func (p *plug) Init(ctx context.Context, cfg *viper.Viper, args []string) error {
	l := logr.FromContextOrDiscard(ctx)
	l.Info("logging the value of my-flag", "value", cfg.GetString("my-flag"))
	l.Info("plugin initialization called")
	return nil
}

func (p *plug) BindFlags(flagset *pflag.FlagSet) *pflag.FlagSet {
	flags.BindTo(flagset)
	return flagset
}

func (p *plug) Version() semver.Version {
	return *vers
}

func (p *plug) ExecuteChecks(ctx context.Context) error {
	l := logr.FromContextOrDiscard(ctx)
	l.Info("Execute Checks Called")
	return nil
}

func (p *plug) Results(ctx context.Context) types.Results {
	return types.Results{}
}

func (p *plug) Submit(ctx context.Context) error {
	l := logr.FromContextOrDiscard(ctx)
	l.Info("Submit called")
	return nil
}
