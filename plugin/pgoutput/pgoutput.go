package pgoutput

import "github.com/piusalfred/pglogrepl/plugin"

var _ plugin.Interface = (*Plugin)(nil)

const pluginName = "pgoutput"

type Plugin struct {
	version int32
}

func (p *Plugin) Options() []string {
	return []string{}
}

func (p *Plugin) Name() string {
	return pluginName
}

func (p *Plugin) Version() int32 {
	return p.version
}
