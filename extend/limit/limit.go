package limit

import (
	"io"

	frpNet "github.com/fatedier/frp/utils/net"
	"golang.org/x/time/rate"
)

const (
	B uint64 = 1 << (10 * (iota))
	KB
	MB
	GB
	TB
	PB
	EB
)

const BurstLimit = 1024 * 1024 * 1024

type Conn struct {
	frpNet.Conn

	lr io.Reader
	lw io.Writer
}

func NewLimitConn(maxread, maxwrite uint64, c frpNet.Conn) Conn {
	return Conn{
		lr:   NewReaderWithLimit(c, maxread),
		lw:   NewWriterWithLimit(c, maxwrite),
		Conn: c,
	}
}

func NewLimitConnWithBucket(c frpNet.Conn, rBucket, wBucket *rate.Limiter) Conn {
	return Conn{
		lr:   NewReaderWithLimitWithBucket(c, rBucket),
		lw:   NewWriterWithLimitWithBucket(c, wBucket),
		Conn: c,
	}
}

func (c Conn) Read(p []byte) (n int, err error) {
	return c.lr.Read(p)
}

func (c Conn) Write(p []byte) (n int, err error) {
	return c.lw.Write(p)
}
