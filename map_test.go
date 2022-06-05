package keepcase_test

import (
	"testing"

	"github.com/aertje/keepcase"
	"github.com/stretchr/testify/assert"
)

func TestNewMapEmpty(t *testing.T) {
	m := keepcase.NewMap[string](nil)

	mm := m.GetBacking()
	assert.Empty(t, mm)
}

func TestNewMapBacking(t *testing.T) {
	mm := make(map[string]string)
	mm["saul"] = "goodman"

	m := keepcase.NewMap(mm)

	// Check that existing entries are processed
	saul, _ := m.Get("saul")
	assert.Equal(t, "goodman", saul)

	// Assert that passed-in map is updated
	m.Set("saul", "GOODMAN")
	assert.Equal(t, "GOODMAN", mm["saul"])
}

func TestSetCaseRespect(t *testing.T) {
	m := keepcase.NewMap[string](nil)

	m.SetCaseRespect("saul", "goodman")

	mm := m.GetBacking()
	assert.Equal(t, "goodman", mm["saul"])
	assert.Equal(t, 1, m.Len())

	m.SetCaseRespect("SAUL", "GOODMAN")

	assert.Equal(t, "GOODMAN", mm["saul"])
	assert.Equal(t, 1, m.Len())
}

func TestSetCaseOverride(t *testing.T) {
	m := keepcase.NewMap[string](nil)

	m.SetCaseOverride("saul", "goodman")

	mm := m.GetBacking()
	assert.Equal(t, "goodman", mm["saul"])
	assert.Equal(t, 1, m.Len())

	m.SetCaseOverride("SAUL", "GOODMAN")

	assert.Equal(t, "GOODMAN", mm["SAUL"])
	assert.Equal(t, 1, m.Len())
}

func TestGetCaseInsensitive(t *testing.T) {
	m := keepcase.NewMap[string](nil)

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
	m := keepcase.NewMap[string](nil)

	m.SetCaseRespect("saul", "goodman")

	saul, _ := m.GetCaseSensitive("saul")
	assert.Equal(t, "goodman", saul)

	saul, ok := m.GetCaseSensitive("SAUL")
	assert.False(t, ok)
	assert.Equal(t, "", saul)

	pollos, ok := m.GetCaseSensitive("pollos")
	assert.False(t, ok)
	assert.Empty(t, "", pollos)
}

func TestSetGet(t *testing.T) {
	m := keepcase.NewMap[string](nil)

	m.Set("saul", "goodman")
	saul, _ := m.Get("saul")

	assert.Equal(t, "goodman", saul)
}
