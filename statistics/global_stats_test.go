package statistics

import (
	"testing"

	"github.com/kyrremann/unparsd/parsing"
	"github.com/stretchr/testify/assert"
)

func TestAllMy(t *testing.T) {
	db, err := parsing.LoadJsonIntoDatabase("../fixture/untappd.json")
	assert.NoError(t, err)

	globalStats, err := AllMyStats(db)
	assert.NoError(t, err)
	assert.Equal(t, 126, globalStats.Checkins)
	assert.Equal(t, 112, globalStats.UniqueBeers)
	assert.Equal(t, "2016-03-01", globalStats.StartDate)
	assert.GreaterOrEqual(t, globalStats.DaysDrinking, 2128)
	assert.LessOrEqual(t, 0.05, globalStats.BeersPerDay)
	assert.Len(t, globalStats.Periods, 5)
	assert.Len(t, globalStats.Periods["2016"].Months, 3)
	assert.Equal(t, "March", globalStats.Periods["2016"].Months[0].Month)
}

func TestDaysInMonth(t *testing.T) {
	days, err := daysInMonth("2016", "02")
	assert.NoError(t, err)
	assert.Equal(t, 29, days)

	days, err = daysInMonth("2017", "02")
	assert.NoError(t, err)
	assert.Equal(t, 28, days)
}

func TestDaysInYear(t *testing.T) {
	days := daysInYear(2016)
	assert.Equal(t, 366, days)

	days = daysInYear(2017)
	assert.Equal(t, 365, days)

	days = daysTillNowInYear()
	assert.Less(t, days, 366)
	assert.Greater(t, days, 0)
}

func TestDaysSince(t *testing.T) {
	days, err := daysSince("1987-07-09")
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, days, 12716)
}

func TestGetMonthAsString(t *testing.T) {
	month, err := getMonthAsString("10")
	assert.NoError(t, err)
	assert.Equal(t, "October", month)
}
