package tests

import (
	"context"
	"database/sql"
	"encoding/json"

	"gorm.io/gorm"

	"github.com/Feresey/scbot/internal/reposiory"
)

func (s *Suite) TestPlayers() {
	r := s.Require()
	ctx := context.Background()
	const playerID = 42

	err := s.cleaner.CleanAll(ctx)
	r.NoError(err)

	players := s.rep.Players(nil)
	s.Run("get empty player", func() {
		player, err := players.GetPlayer(ctx, playerID)
		r.Nil(player)
		r.ErrorIs(err, gorm.ErrRecordNotFound)
	})

	s.Run("create player", func() {
		player := &reposiory.Player{
			UserID:  playerID,
			Comment: "returd",
		}
		err := players.CreatePlayer(ctx, player)
		r.NoError(err)

		player, err = players.GetPlayer(ctx, playerID)
		r.NoError(err)
		wantPlayer := &reposiory.Player{
			UserID:    playerID,
			Comment:   "returd",
			CreatedAt: player.CreatedAt,
		}
		r.Equal(wantPlayer, player)
	})
}

func (s *Suite) TestCorporations() {
	r := s.Require()
	ctx := context.Background()
	const corpID = 69

	corps := s.rep.Corporations(nil)
	s.Run("get empty corp", func() {
		corp, err := corps.GetCorp(ctx, corpID)
		r.ErrorIs(err, gorm.ErrRecordNotFound)
		r.Nil(corp)
	})

	s.Run("create corp", func() {
		corp := &reposiory.Corporation{
			CorpID: corpID,
		}
		err := corps.CreateCorp(ctx, corp)
		r.NoError(err)

		corp, err = corps.GetCorp(ctx, corpID)
		r.NoError(err)
		wantcorp := &reposiory.Corporation{
			CorpID:    corpID,
			CreatedAt: corp.CreatedAt,
		}
		r.Equal(wantcorp, corp)
	})
}

func (s *Suite) TestRepository() {
	r := s.Require()
	ctx := context.Background()
	const (
		playerID     = 42
		corpID       = 69
		secondCorpID = 420
	)
	s.Run("add player info for new player & corp", func() {
		err := s.rep.AddPlayerInfo(ctx, reposiory.PlayerInfo{
			UserID:   playerID,
			CorpID:   sql.NullInt32{Int32: corpID},
			Nickname: "nick",
			Info:     json.RawMessage(`{"name": "nick"}`),
		})
		r.NoError(err)
	})

	s.Run("player changed corp", func() {
		err := s.rep.AddPlayerInfo(ctx, reposiory.PlayerInfo{
			UserID:   playerID,
			CorpID:   sql.NullInt32{Int32: secondCorpID},
			Nickname: "nick",
			Info:     json.RawMessage(`{"name": "nick"}`),
		})
		r.NoError(err)
	})

	s.Run("player changed nick", func() {
		err := s.rep.AddPlayerInfo(ctx, reposiory.PlayerInfo{
			UserID:   playerID,
			CorpID:   sql.NullInt32{Int32: secondCorpID},
			Nickname: "nickname",
			Info:     json.RawMessage(`{"name": "nick"}`),
		})
		r.NoError(err)
	})

	s.Run("player leaves corp", func() {
		err := s.rep.AddPlayerInfo(ctx, reposiory.PlayerInfo{
			UserID:   playerID,
			Nickname: "nickname",
			Info:     json.RawMessage(`{"name": "nick"}`),
		})
		r.NoError(err)
	})

	s.Run("get player history", func() {
		history, err := s.rep.GetPlayerHistory(ctx, playerID)
		r.NoError(err)
		r.NotNil(history)

		r.Len(history.History, 4)
	})
}
