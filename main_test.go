package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchAuthorBooks_Success(t *testing.T) {
	db, err := initDB()
	require.NoError(t, err)
	defer db.Close()

	author := Author{ID: 1}
	books, found, err := SearchAuthorBooks(db, author)

	require.NoError(t, err)
	require.True(t, found)
	require.NotEmpty(t, books)
}

func TestSearchAuthorBooks_Fail(t *testing.T) {
	db, err := initDB()
	require.NoError(t, err)
	defer db.Close()

	author := Author{ID: 999}
	books, found, err := SearchAuthorBooks(db, author)

	require.NoError(t, err)
	require.False(t, found)
	require.Empty(t, books)
}

func TestSearchAuthorBooks_EdgeCase(t *testing.T) {
	db, err := initDB()
	require.NoError(t, err)
	defer db.Close()

	author := Author{ID: -1}
	books, found, err := SearchAuthorBooks(db, author)

	require.NoError(t, err)
	require.False(t, found)
	require.Empty(t, books)
}
