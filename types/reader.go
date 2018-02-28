package types

import "time"

type ConnReader interface {
	RefreshReadTimeout(d time.Duration)}
