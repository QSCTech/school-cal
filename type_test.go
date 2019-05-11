package school_cal

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
	// check time parsing
	assert.True(
		t,
		schoolYear.Start.Equal(
			time.Date(2015, time.August, 10, 0, 0, 0, 0, local),
		),
	)

	// check semesters
	AW := schoolYear.Semesters.AW
	assert.NotNil(t, AW)
	assert.False(t, AW.StartsWithWeekZero)

	// check holidays
	assert.Equal(t, 1, len(schoolYear.Holidays))
	holiday := schoolYear.Holidays[0]
	assert.Equal(t, "春学期考试周", holiday.Name)
}
