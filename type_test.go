package schoolcal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func TestSchoolYearsUnmarshal(t *testing.T) {
	SampleCal, err := ioutil.ReadFile("sample_cal.json")
	assert.Nil(t, err)
	schoolYears := make(map[string]*SchoolYear)
	assert.Nil(t, json.Unmarshal(SampleCal, &schoolYears))
	assert.Equal(t, 1, len(schoolYears))
	schoolYear := schoolYears["2015-2016"]
	assert.NotNil(t, schoolYear)

	local, err := time.LoadLocation("Asia/Chongqing")
	assert.Nil(t, err)

	assert.True(
		t,
		schoolYear.Start.Equal(
			time.Date(2015, time.August, 10, 0, 0, 0, 0, local),
		),
	)

	assert.True(
		t,
		schoolYear.End.Equal(
			time.Date(2016, time.August, 1, 0, 0, 0, 0, local),
		),
	)

	// check semesters
	AW := schoolYear.Semesters.AW
	assert.NotNil(t, AW)
	assert.False(t, AW.StartsWithWeekZero)
	assert.True(
		t,
		AW.Start.Equal(
			time.Date(2015, time.September, 14, 0, 0, 0, 0, local),
		),
	)

	assert.True(
		t,
		AW.End.Equal(
			time.Date(2016, time.January, 24, 0, 0, 0, 0, local),
		),
	)

	// check holidays
	assert.Equal(t, 1, len(schoolYear.Holidays))
	holiday := schoolYear.Holidays[0]
	assert.Equal(t, "春学期考试周", holiday.Name)

	assert.True(
		t,
		holiday.Start.Equal(
			time.Date(2016, time.April, 23, 0, 0, 0, 0, local),
		),
	)

	assert.True(
		t,
		holiday.End.Equal(
			time.Date(2016, time.April, 28, 0, 0, 0, 0, local),
		),
	)

	// check adjustments
	assert.Equal(t, 2, len(schoolYear.Adjustments))
	adjustment := schoolYear.Adjustments[0]
	assert.Equal(t, "0", adjustment.Term)
	assert.Equal(t, "国庆节", adjustment.Name)
	assert.True(
		t,
		adjustment.FromStart.Equal(
			time.Date(2015, time.October, 3, 0, 0, 0, 0, local),
		),
	)

	assert.True(
		t,
		adjustment.FromEnd.Equal(
			time.Date(2015, time.October, 5, 0, 0, 0, 0, local),
		),
	)
	assert.True(
		t,
		adjustment.ToStart.Equal(
			time.Date(2015, time.October, 6, 0, 0, 0, 0, local),
		),
	)

	assert.True(
		t,
		adjustment.ToEnd.Equal(
			time.Date(2015, time.October, 8, 0, 0, 0, 0, local),
		),
	)

	// check special
	assert.Equal(t, 2, len(schoolYear.Specials))
	special := schoolYear.Specials[0]
	assert.Equal(t, "371E0010", special.Code)
	assert.Equal(t, "odd", special.Weekly)
	assert.True(
		t,
		special.Date.Equal(
			time.Date(2016, time.September, 12, 0, 0, 0, 0, local),
		),
	)
}
