package add_source_name

import (
	"github.com/ozontech/file.d/fd"
	"github.com/ozontech/file.d/pipeline"
	"github.com/ozontech/file.d/plugin"
)

/*{ introduction
It adds field containing source name to an event.
}*/

type Plugin struct {
	config *Config
	plugin.NoMetricsPlugin
}

// ! config-params
// ^ config-params
type Config struct {
	// > @3@4@5@6
	// >
	// > The event field to which put the source name. Must be a string.
	Field string `json:"field" default:"source_name" required:"true"` // *
}

func init() {
	fd.DefaultPluginRegistry.RegisterAction(&pipeline.PluginStaticInfo{
		Type:    "add_source_name",
		Factory: factory,
	})
}

func factory() (pipeline.AnyPlugin, pipeline.AnyConfig) {
	return &Plugin{}, &Config{}
}

func (p *Plugin) Start(config pipeline.AnyConfig, _ *pipeline.ActionPluginParams) {
	p.config = config.(*Config)
}

func (p *Plugin) Stop() {
}

func (p *Plugin) Do(event *pipeline.Event) pipeline.ActionResult {
	event.Root.AddFieldNoAlloc(event.Root, p.config.Field).MutateToString(event.SourceName)
	return pipeline.ActionPass
}
