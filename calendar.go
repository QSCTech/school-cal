package schoolcal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	// DefaultAPIURL is the default url of school calendar API
	DefaultAPIURL = "https://product.zjuqsc.com/schoolCal/backend-api/output/"

	// DefaultCacheDuration is the default duration of calendar data cache
	DefaultCacheDuration = time.Hour * 24
)

// DefaultInitErrorHandler is the default error handler when calendar init;
func DefaultInitErrorHandler(err error) {
	log.Fatalf("init error: %s\n", err.Error())
}

// DefaultUpdateErrorHandler is the default error handler when calendar update;
func DefaultUpdateErrorHandler(err error) {
	log.Printf("update error: %s\n", err.Error())
}

// Calendar is a interface to control school years
type Calendar interface {
	// GetSchoolYears is a method to get school years
	GetSchoolYears() map[string]*SchoolYear
}

// CalendarOptions is a struct to construct calendar
type CalendarOptions struct {
	// APIURL is the URL of school calendar API
	APIURL string

	// CacheDuration is the duration of calendar data cache
	CacheDuration time.Duration

	// InitErrorHandler is the error handler when calendar init
	// should exit when some error occurred, otherwise you will get nil school years
	InitErrorHandler func(error)

	// UpdateErrorHandler is the error handler when calendar update
	// should just log the error
	UpdateErrorHandler func(error)
}

type calendar struct {
	apiAddr            string
	cacheDuration      time.Duration
	initErrorHandler   func(error)
	updateErrorHandler func(error)
	schoolYears        atomic.Value // map[string]*SchoolYear
}

func newCalendar() *calendar {
	return &calendar{
		apiAddr:            DefaultAPIURL,
		cacheDuration:      DefaultCacheDuration,
		initErrorHandler:   DefaultInitErrorHandler,
		updateErrorHandler: DefaultUpdateErrorHandler,
	}
}

func (cal *calendar) init() {
	cal.update(cal.initErrorHandler)
	ticker := time.NewTicker(cal.cacheDuration)
	go func() {
		for range ticker.C {
			cal.update(cal.updateErrorHandler)
		}
	}()
}

func (cal *calendar) update(errHandler func(err error)) {
	resp, err := http.Get(cal.apiAddr)
	if err != nil {
		errHandler(err)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			errHandler(err)
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errHandler(err)
		return
	}

	schoolYears := make(map[string]*SchoolYear)
	if err = json.Unmarshal(data, &schoolYears); err != nil {
		errHandler(err)
		return
	}
	cal.schoolYears.Store(schoolYears)
}

// GetSchoolYears is a method to get school years
func (cal *calendar) GetSchoolYears() map[string]*SchoolYear {
	return cal.schoolYears.Load().(map[string]*SchoolYear)
}

// NewCalendar is a method to construct a new Calender
func NewCalendar(option *CalendarOptions) Calendar {
	calendar := newCalendar()
	if option != nil {
		if option.APIURL != "" {
			calendar.apiAddr = option.APIURL
		}

		if option.CacheDuration != 0 {
			calendar.cacheDuration = option.CacheDuration
		}

		if option.InitErrorHandler != nil {
			calendar.initErrorHandler = option.InitErrorHandler
		}

		if option.UpdateErrorHandler != nil {
			calendar.updateErrorHandler = option.UpdateErrorHandler
		}
	}

	calendar.init()
	return calendar
}
