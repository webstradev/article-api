package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllArticles(t *testing.T) {
	got := getAllArticles()
	require.ElementsMatch(t, got, articleList)
}
