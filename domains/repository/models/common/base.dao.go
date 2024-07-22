package common

import "time"

type CreateBaseDAO struct {
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

func (b *CreateBaseDAO) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = currentDate
	}
	b.UpdatedAt = currentDate
}

// ------------------------------

type GetBaseDAO struct {
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

// ------------------------------

type UpdateBaseDAO struct {
	UpdatedAt time.Time  `bson:"updated_at"`
}

func (u *UpdateBaseDAO) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
