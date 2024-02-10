package config

import (
	"time"
)

// These are hardcoded right now, but may be options eventually
var (
	// Server
	WEB_PORT = 9098

	// Update Interval
	INTERVAL = time.Duration(5)
)
