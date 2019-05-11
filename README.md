[![Go Report Card](https://goreportcard.com/badge/github.com/QSCTech/school-cal)](https://goreportcard.com/report/github.com/QSCTech/school-cal)
[![Build Status](https://travis-ci.org/QSCTech/school-cal.svg?branch=master)](https://travis-ci.org/QSCTech/school-cal)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/QSCTech/school-cal/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/QSCTech/school-cal?status.svg)](https://godoc.org/github.com/QSCTech/school-cal)

### Example

using go mod

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/QSCTech/school-cal"
)

func main() {
	cal := schoolcal.NewCalendar(nil)
	data, _ := json.Marshal(cal.GetSchoolYears())
	fmt.Println(string(data))
}
```

### Options

You can configure Calendar by delivering [CalendarOption](https://godoc.org/github.com/QSCTech/school-cal#CalendarOptions) into NewCalendar