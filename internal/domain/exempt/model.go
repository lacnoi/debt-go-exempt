package exempt

import "time"

type Exempt struct {
	CANo         string     `json:"ca_no"`
	BANo         string     `json:"ba_no"`
	MobileNum    string     `json:"mobile_num"`
	ModeID       string     `json:"mode_id"`
	EffectiveDat time.Time  `json:"effective_dat"`
	EndDat       *time.Time `json:"end_dat,omitempty"`
	Created      time.Time  `json:"created"`
	CreatedBy    string     `json:"created_by"`
	LastUpd      *time.Time `json:"last_upd,omitempty"`
	LastUpdBy    *string    `json:"last_upd_by,omitempty"`
}
