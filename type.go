package schoolcal

// SchoolYear is struct to describe a school year in school calendar
type SchoolYear struct {
	Start CustomTime `json:"start"`

	// 到 end 为止，不包括 end 代表的这一天，即左闭右开区间的惯例。
	End         CustomTime    `json:"end"`
	Semesters   Semesters     `json:"semesters"`
	Holidays    []*Holiday    `json:"holidays"`
	Adjustments []*Adjustment `json:"adjustments"`
	Specials    []*Special    `json:"special"`
}

// Semesters is struct to describe all the semester in a school year
type Semesters struct {
	// SM：Summer Mini，夏季短学期
	SM Semester `json:"SM"`
	// AW：Autumn-Winter，秋冬长学期
	AW Semester `json:"AW"`
	// Au：Autumn，秋季学期
	Au Semester `json:"Au"`
	// Wi：Winter，冬季学期
	Wi Semester `json:"Wi"`
	// WM：Winter Mini，冬季短学期（寒假）
	WM Semester `json:"WM"`
	// SS：Spring-Summer，春夏学期
	SS Semester `json:"SS"`
	// Sp：Spring，春季学期
	Sp Semester `json:"Sp"`
	// Su：Summer，夏季学期
	Su Semester `json:"Su"`
	// ST：Summer Tiny，夏季小学期
	ST Semester `json:"ST"`
}

// Semester is a struct to describe single semester
type Semester struct {
	Start              CustomTime `json:"start"`
	End                CustomTime `json:"end"`
	StartsWithWeekZero bool       `json:"startsWithWeekZero"`
}

// Holiday is a struct to describe single holiday
type Holiday struct {
	Name  string     `json:"name"`
	Start CustomTime `json:"start"`
	End   CustomTime `json:"end"`
}

// Adjustment is a struct to describe single adjustment
type Adjustment struct {
	//"term" 字段表示需要调整的课程，可针对小学期调课，“无筛选” 为兼容以前的调课逻辑，不对课程学期进行筛选。
	// <option value="0">无筛选</option>
	// <option value="1">秋冬</option>
	// <option value="2">秋</option>
	// <option value="3">冬</option>
	// <option value="5">春夏</option>
	// <option value="6">春</option>
	// <option value="7">夏</option>
	Term string `json:"term"`

	Name string `json:"name"`

	// 左闭右开，10-3~4 => 10-6~7。
	FromStart CustomTime `json:"fromStart"`

	FromEnd CustomTime `json:"fromEnd"`

	ToStart CustomTime `json:"toStart"`

	ToEnd CustomTime `json:"toEnd"`
}

// Special is a struct to describe single special
// 特殊课程，比如行策
type Special struct {
	Code string `json:"code"`
	// 课程周期 all:每周，odd：单周，even：双周
	Weekly string     `json:"weekly"`
	Date   CustomTime `json:"date"`
}
