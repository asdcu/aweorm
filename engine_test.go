package main

import (
	"aweorm/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEngine_NewSession(t *testing.T) {
	log.SetLevel(log.Disabled)
	engine, err := NewEngine("sqlite3", "awe.db")
	assert.NoError(t, err)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, err = s.Raw("CREATE TABLE User(Name text);").Exec()
	assert.Error(t, err)
	result, err := s.Raw("INSERT INTO User(`Name`) Values (?), (?)", "Tome", "Sam").Exec()
	assert.NoError(t, err)
	count, err := result.RowsAffected()
	assert.NoError(t, err)
	t.Log(count)
	lid, _ := result.LastInsertId()
	t.Log(lid)
}
