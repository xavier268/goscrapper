package rt

import "time"

// Default duration when SLOW is requested with invalid millis expression.
// Can be adjusted before running a request calling SLOW.
// Use for debugging only.
var SLOW_DELAY = 1 * time.Second
