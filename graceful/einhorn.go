// +build !windows

package graceful

import (
	"log"
	"os"
	"strconv"
	"syscall"
)

func init() {
	// This is a little unfortunate: gopi/bind already knows whether we're
	// running under einhorn, but we don't want to introduce a dependency
	// between the two packages. Since the check is short enough, inlining
	// it here seems "fine."
	mpid, err := strconv.Atoi(os.Getenv("EINHORN_MASTER_PID"))
	if err != nil || mpid != os.Getppid() {
		return
	}

	log.Print("graceful: Einhorn detected, adding SIGUSR2 handler")
	AddSignal(syscall.SIGUSR2)
}
