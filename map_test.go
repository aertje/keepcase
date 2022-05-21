package keepcase_test

import (
	"testing"

	"github.com/aertje/keepcase"
	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	m := keepcase.NewMap[string]()

	mm := m.AsMap()
	assert.Empty(t, mm)
}

func TestSetCaseRespect(t *testing.T) {
	m := keepcase.NewMap[string]()

	m.SetCaseRespect("saul", "goodman")

	mm := m.AsMap()
	assert.Equal(t, "goodman", mm["saul"])
	assert.Equal(t, 1, m.Len())

	m.SetCaseRespect("SAUL", "GOODMAN")

	mm = m.AsMap()

	assert.Equal(t, "GOODMAN", mm["saul"])
	assert.Equal(t, 1, m.Len())
}

func TestSetCaseOverride(t *testing.T) {
	m := keepcase.NewMap[string]()

	m.SetCaseOverride("saul", "goodman")

	mm := m.AsMap()
	assert.Equal(t, "goodman", mm["saul"])
	assert.Equal(t, 1, m.Len())

	m.SetCaseOverride("SAUL", "GOODMAN")

	mm = m.AsMap()

	assert.Equal(t, "GOODMAN", mm["SAUL"])
	assert.Equal(t, 1, m.Len())
}

func TestGetCaseInsensitive(t *testing.T) {
	m := keepcase.NewMap[string]()

	m.SetCaseRespect("saul", "goodman")

	saul, _ := m.GetCaseInsensitive("saul")
	assert.Equal(t, "goodman", saul)

	saul, _ = m.GetCaseInsensitive("SAUL")
	assert.Equal(t, "goodman", saul)

	pollos, ok := m.GetCaseInsensitive("pollos")
	assert.False(t, ok)
	assert.Equal(t, "", pollos)
}

func TestGetCaseSensitive(t *testing.T) {
	m := keepcase.NewMap[string]()

	m.SetCaseRespect("saul", "goodman")

	saul, _ := m.GetCaseSensitive("saul")
	assert.Equal(t, "goodman", saul)

	saul, ok := m.GetCaseSensitive("SAUL")
	assert.False(t, ok)
	assert.Equal(t, "", saul)

	pollos, ok := m.GetCaseSensitive("pollos")
	assert.False(t, ok)
	assert.Equal(t, "", pollos)
}

func TestSetGet(t *testing.T) {
	m := keepcase.NewMap[string]()

	m.Set("saul", "goodman")
	saul, _ := m.Get("saul")

	assert.Equal(t, "goodman", saul)
}
