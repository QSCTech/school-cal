package schoolcal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalendar(t *testing.T) {
	calendar := NewCalendar(nil)
	schoolYears := calendar.GetSchoolYears()
	assert.NotNil(t, schoolYears)
	assert.True(t, len(schoolYears) > 0)
}
