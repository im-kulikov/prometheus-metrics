package prometheus

import (
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Serve prometheus metrics
func Serve(l Logger) {
	if l == nil {
		l = defaultLogger
	}

	bind := os.Getenv("PROMETHEUS")

	if len(bind) == 0 {
		return
	}

	l.Debugf("run prometheus metrics on `%s`", bind)

	go func() {
		if err := http.ListenAndServe(bind, promhttp.Handler()); err != nil {
			l.Panic(errors.Wrap(err, "can't serve prometheus metrics"))
		}
	}()
}
