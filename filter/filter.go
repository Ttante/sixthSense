package filter

import (
	"github.com/iovisor/gobpf"
)

/*
#cgo CFLAGS: -I/usr/include/bcc/compat
#cgo LDFLAGS: -lbcc
#include <bcc/bpf_common.h>
#include <bcc/libbpf.h>
void perf_reader_free(void *ptr);
*/
import "C"

const packetFilter = `

`

type Filter struct {
	SRC string
}
