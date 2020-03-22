package main

import (
	"github.com/matei207/github-issue-triage/pkg/client"
	"github.com/matei207/github-issue-triage/pkg/helper"
	"github.com/matei207/github-issue-triage/pkg/options"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.SugaredLogger

func init() {
	logger = helper.NewZapSugaredLogger()
}

func main() {
	defer logger.Sync()

	opt := &options.Options{}
	opt.AddFlags(pflag.CommandLine)
	opt.RegisterOptions()
	pflag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := []byte("Hello, World!")
		if _, err := w.Write(b); err != nil {
			logger.Infof("error sending resp:%v", err)
		}
	})

	client.Run(opt)

	if err := Run(opt); err != nil {
		logger.Infof("error listening on :8080")
	}

}

func Run(opt *options.Options) error {
	return http.ListenAndServe(opt.ListenPort, nil)
}
