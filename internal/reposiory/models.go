package reposiory

import (
	"time"

	"github.com/Feresey/scbot/internal/infogetter"
)

type Corporation struct {
	CorpID    int `gorm:"primary_key"`
	CreatedAt time.Time
}

type CorporationInfo struct {
	ID        int `gorm:"primary_key"`
	CorpID    int `gorm:"references:corporation.corp_id"`
	CreatedAt time.Time
}

type Player struct {
	UserID    int `gorm:"primary_key"`
	CorpID    int
	CreatedAt time.Time
}

type PlayerInfo struct {
	ID        int `gorm:"primary_key"`
	UserID    int `gorm:"references:player.user_id"`
	CorpID    int `gorm:"references:corporation.corp_id"`
	NickName  string
	Info      infogetter.PlayerInfo `gorm:"type:jsonb"`
	CreatedAt time.Time
}

// type UserInfo struct {
// 	ID          int
// 	Nickname    string
// 	Corporation *Corporation
// 	Comment     string
// 	Info        infogetter.PlayerInfo
// 	UpdatedAt   time.Time
// }

// type UserHistory struct {
// 	Corporation *Corporation
// }

// type Corporation struct {
// 	ID        int
// 	Name      string
// 	Tag       string
// 	UpdatedAt time.Time
// }

// type CorporationMembers struct {
// 	CorpInfo  Corporation
// 	UserNames []string
// }
