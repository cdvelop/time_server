package timeserver

import "time"

func (TimeServer) UnixNano() int64 {
	return time.Now().UnixNano()
}
