package options

import flag "github.com/spf13/pflag"

const (
	defaultListenPort = ":8080"
)

var Opt *Options

type Options struct {
	ListenPort string
	ApiToken string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	flag.StringVar(&o.ListenPort, "port", defaultListenPort, "HTTP Port to listen on")
	flag.StringVar(&o.ApiToken, "token", o.ApiToken, "API token for GitHub requests")
}

func (o *Options) RegisterOptions() {
	Opt = o
}
