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
	DefaultAPIURL        = "https://product.zjuqsc.com/schoolCal/backend-api/output/"
	DefaultCacheDuration = time.Hour * 24
)

func DefaultInitErrorHandler(err error) {
	log.Fatalf("init error: %s\n", err.Error())
}

func DefaultUpdateErrorHandler(err error) {
	log.Printf("update error: %s\n", err.Error())
}

type Calendar interface {
	GetSchoolYears() map[string]*SchoolYear
}

type CalendarOptions struct {
	APIURL             string
	CacheDuration      time.Duration
	InitErrorHandler   func(error)
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

func (cal *calendar) GetSchoolYears() map[string]*SchoolYear {
	return cal.schoolYears.Load().(map[string]*SchoolYear)
}

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
