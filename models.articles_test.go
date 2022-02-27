package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllArticles(t *testing.T) {
	got := getAllArticles()
	require.ElementsMatch(t, got, articleList)
}

func TestGetArticleByID(t *testing.T) {
	tests := []struct {
		name   string
		id     int
		exp    *article
		expErr error
	}{
		{"existing article", 1, &article{ID: 1, Title: "Article 1", Content: "Article 1 body"}, nil},
		{"non existing article", 3, nil, errors.New("article not found")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := getArticleByID(test.id)
			if err != nil {
				// Check that error matches
				require.Equal(t, err, test.expErr)
			} else {
				require.Nil(t, test.expErr)
			}

			// Check that the article matches
			require.Equal(t, got, test.exp)
		})
	}

}
