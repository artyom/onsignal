// Package onsignal provides helper functions to set up signal-handling
// functions.
//
// Helpers provided by this package are mainly intended to be used in top-level
// main() functions.
package onsignal

import (
	"os"
	"os/signal"
)

// Handler type functions can be registered by Once and Repeat calls. Handler is
// called with received signal as its argument.
type Handler func(os.Signal)

// Once schedules fn to be called one time on any of sig signals.
func Once(fn Handler, sig ...os.Signal) {
	if len(sig) == 0 {
		return
	}
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, sig...)
		defer signal.Stop(sigCh)
		fn(<-sigCh)
	}()
}

// Repeat schedules fn to be called for each of sig signals received.
func Repeat(fn Handler, sig ...os.Signal) {
	if len(sig) == 0 {
		return
	}
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, sig...)
		defer signal.Stop(sigCh)
		for {
			fn(<-sigCh)
		}
	}()
}
