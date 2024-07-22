package dao

import "time"

type UpdateFollowerDAOReq struct {
	Commentary string     `bson:"commentary,omitempty"`
	Score      float32    `bson:"score,omitempty"`
	Anonymous  bool       `bson:"anonymous,omitempty"`
	UpdatedAt  time.Time  `bson:"updated_at,omitempty"`
}

func (u *UpdateFollowerDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
