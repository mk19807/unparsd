package statistics

import (
	"testing"

	"github.com/kyrremann/unparsd/parsing"
	"github.com/stretchr/testify/assert"
)

func TestBeers(t *testing.T) {
	db, err := parsing.LoadJsonIntoDatabase("../fixture/untappd.json")
	assert.NoError(t, err)

	beers, err := Beers(db)
	assert.NoError(t, err)
	assert.Len(t, beers, 82)

	checkins := 0
	for _, beer := range beers {
		checkins += beer.Checkins
	}
	assert.Equal(t, 100, checkins)
}