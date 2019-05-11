package schoolcal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Task struct {
	Date CustomTime `json:"date"`
}

const TaskJSON = `{"date":"2015-08-10T00:00+08:00"}`

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	task := new(Task)
	assert.Nil(t, json.Unmarshal([]byte(TaskJSON), task))
	date := task.Date
	assert.Equal(t, 2015, date.Year())
	assert.Equal(t, time.August, date.Month())
	assert.Equal(t, 10, date.Day())
	assert.Equal(t, 0, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())
}

func TestCustomTime_MarshalJSON(t *testing.T) {
	task := new(Task)
	assert.Nil(t, json.Unmarshal([]byte(TaskJSON), task))
	data, err := json.Marshal(task)
	assert.Nil(t, err)
	assert.Equal(t, TaskJSON, string(data))
}
