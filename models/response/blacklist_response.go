package response

import "time"

type ShowBlackListDto struct {
	Items map[string]time.Time
}
