package database

import (
	"context"
	"log"
	"strconv"
	"testing"

	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	"github.com/NeverlandMJ/GRUD/grud-service/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	t.Run("delete post", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		data = append(data, generateFakeData(5, 1, 1)...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		err = p.DeletePost(context.Background(), 1)
		require.NoError(t, err)

		got, err := p.GetPostByID(context.Background(), 1)
		require.ErrorIs(t, err, errs.ErrPostNotFound)
		assert.Empty(t, got)
	})

	t.Run("update title", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		data = append(data, generateFakeData(1, 1, 1)...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		newTitle := "I am a new title"
		got, err := p.UpdateTitle(context.Background(), 1, newTitle)
		require.NoError(t, err)
		assert.Equal(t, newTitle, got.Title)
	})
	t.Run("update body", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		data = append(data, generateFakeData(1, 1, 1)...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		newBody := "I am a new body"
		got, err := p.UpdateBody(context.Background(), 1, newBody)
		require.NoError(t, err)
		assert.Equal(t, newBody, got.Body)
	})
	t.Run("list posts", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		data := make([]entity.Data, 0)
		d := generateFakeData(50, 1, 1)
		data = append(data, d...)
		err := p.createPost(context.Background(), data)
		require.NoError(t, err)

		want1 := d[:5]
		got1, count, err := p.ListPosts(context.Background(), 1, 5)
		require.NoError(t, err)
		assert.Equal(t, len(data), count)
		assert.Equal(t, want1, got1)

		want2 := d[5:10]
		got2, count, err := p.ListPosts(context.Background(), 2, 5)
		require.NoError(t, err)
		assert.Equal(t, len(data), count)
		assert.Equal(t, want2, got2)
	})
}
