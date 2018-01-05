/*
Package signal manages signal handling.
*/
package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Ignore sets all signals to be ignored
func Ignore(f func()) {
	signal.Ignore()
}

// Hup sets f() to be run on receipt of SIGHUP
// SIGHUP is meant to be used to refresh program configuration
func Hup(f func()) {

	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP)

	go func() {
		for {
			<-sighup
			f()
		}
	}()
}

// Int sets f() to be run on receipt of SIGINT
// SIGINT is meant to be used to stop the program
func Int(f func()) {

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)

	go func() {
		<-sigint
		f()
	}()
}

// Quit sets f() to be run on receipt of SIGQUIT
// SIGQUIT is meant to be used to stop the program gracefully
func Quit(f func()) {

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, syscall.SIGQUIT)

	go func() {
		<-sigquit
		f()
	}()

}

// Term sets f() to be run on receipt of SIGTERM
// SIGTERM is meant to be used to stop the program
func Term(f func()) {

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		<-sigterm
		f()
	}()
}
