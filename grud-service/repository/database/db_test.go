package database

import (
	"context"
	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"strconv"
	"testing"
)

func newTestConfig() config.Config {
	return config.Config{
		Host: "",
		Port: "",
		PostgresConfig: config.PostgresConfig{
			PostgresHost:           "localhost",
			PostgresPort:           "5432",
			PostgresUser:           "sunbula",
			PostgresPassword:       "2307",
			PostgresDB:             "postpb_test",
			PostgresMigrationsPath: "./../postgres/migrations",
		},
	}

}

func generateFakeData(n int, userID int, lastPostID int) []entity.Data {
	datas := make([]entity.Data, 0, n)
	for i := lastPostID; i <= n; i++ {
		data := entity.Data{
			ID:     i,
			UserID: userID,
			Title:  "title" + strconv.Itoa(i),
			Body:   "body" + strconv.Itoa(i),
		}
		datas = append(datas, data)
	}

	return datas
}

func cleanup(p *PostgresDB) func() {
	return func() {
		if err := p.cleanup(context.Background()); err != nil {
			log.Panicln("failed to cleanup db, should be done manually", err)
		}
	}
}

func TestPostgresDB(t *testing.T) {
	p, err := NewPostgresDB(newTestConfig())
	require.NoError(t, err)
	t.Cleanup(cleanup(p))

	t.Run("get posts by userID", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		user1 := append(data, generateFakeData(3, 1, 1)...)
		user2 := append(data, generateFakeData(3, 2, 4)...)
		data = append(data, user1...)
		data = append(data, user2...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		// Testing GetPostByUserID
		got, err := p.GetPostsByUserID(context.Background(), 1)
		require.NoError(t, err)
		assert.Equal(t, user1, got)
	})

	t.Run("get post by postID", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		user1 := append(data, generateFakeData(3, 1, 1)...)
		user2 := append(data, generateFakeData(3, 2, 4)...)
		data = append(data, user1...)
		data = append(data, user2...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		got, err := p.GetPostByID(context.Background(), 1)
		require.NoError(t, err)
		assert.Equal(t, data[0], got)
	})
}
