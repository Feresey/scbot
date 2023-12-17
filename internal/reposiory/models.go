package reposiory

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Corporation struct {
	CorpID    int `gorm:"primary_key"`
	CreatedAt time.Time
}

type CorporationInfo struct {
	ID        int           `gorm:"primary_key"`
	CorpID    sql.NullInt32 `gorm:"references:corporation.corp_id"`
	CreatedAt time.Time
}

type Player struct {
	UserID    int `gorm:"primary_key"`
	Comment   string
	CreatedAt time.Time
}

type PlayerInfo struct {
	ID        int           `gorm:"primary_key"`
	UserID    int           `gorm:"references:player.user_id"`
	CorpID    sql.NullInt32 `gorm:"references:corporation.corp_id"`
	Nickname  string
	Info      json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time
}

type PlayerHistoryItem struct {
	Nickname        string
	Info            json.RawMessage
	CorporationInfo *CorporationInfo
}

type PlayerHistory struct {
	History []*PlayerHistoryItem
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
