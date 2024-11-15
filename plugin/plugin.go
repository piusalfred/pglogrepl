package plugin

type Interface interface {
	Name() string
	Version() int32
	Options() []string
}
