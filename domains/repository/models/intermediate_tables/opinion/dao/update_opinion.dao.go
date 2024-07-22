package dao

import "time"

type UpdateOpinionDAOReq struct {
	Commentary string     `bson:"commentary,omitempty"`
	Score      float32    `bson:"score,omitempty"`
	Anonymous  bool       `bson:"anonymous,omitempty"`
	UpdatedAt  time.Time  `bson:"updated_at,omitempty"`
}

func (u *UpdateOpinionDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
