package timeserver

import "time"

func (timeServer) UnixNano() int64 {
	return time.Now().UnixNano()
}
