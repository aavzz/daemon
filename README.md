# daemon
Daemon initialization routines (process, log, pidfile, signal handling)

```go
import (
	"github.com/aavzz/daemon/fork"
	"github.com/aavzz/daemon/log"
	"github.com/aavzz/daemon/pid"
	"github.com/aavzz/daemon/signal"
}
     
//Initialize logging to syslog with a relevant tag
//The process exits if local syslog is not running
log.InitSyslog("tag")

//Daemonize process if needed
fork.Daemonize()

//After daemonize() this part runs in child only

//Check for double invocation and write pidfile
pid.Write("/path/to/pidfile")

//Set all signals to be ignored
signal.Ignore()

//Setup SIGHUP to re-read configuration
signal.Hup(func() {
	log.Info("SIGHUP received, re-reading configuration file")
	if err := viper.ReadInConfig(); err != nil {
		pid.Remove()
		log.Fatal(err.Error())
	}
})

//Setup SIGTERM to remove pidfile and exit immidiately
signal.Term(func() {
	log.Info("SIGTERM received, exitting")
	pid.Remove()
	os.Exit(0)
})

//signal.Quit is set up to terminate the process gracefully
//in approppriate place (where required objects are available).


//further initialization and an event loop

```
