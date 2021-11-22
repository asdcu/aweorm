package schema

import (
	"aweorm/dialect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	type User struct {
		Name string `aweorm:"PRIMARY KEY"`
		Age  int
	}

	var TestDail, _ = dialect.GetDialect("sqlite3")

	schema := Parse(&User{}, TestDail)
	assert.Equal(t, "User", schema.Name)
	assert.Len(t, schema.Fields, 2)
	assert.Equal(t, "PRIMARY KEY", schema.GetField("Name").Tag)
}
