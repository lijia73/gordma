// +build linux

package ibverbs

//#include <infiniband/verbs.h>
import "C"

// access flag
const (
	IBV_ACCESS_LOCAL_WRITE = C.IBV_ACCESS_LOCAL_WRITE
	IBV_ACCESS_REMOTE_WRITE = C.IBV_ACCESS_REMOTE_WRITE
	IBV_ACCESS_REMOTE_READ = C.IBV_ACCESS_REMOTE_READ
	IBV_ACCESS_REMOTE_ATOMIC = C.IBV_ACCESS_REMOTE_ATOMIC
)

// qp type
const (
	IBV_QPT_RC = C.IBV_QPT_RC
)
