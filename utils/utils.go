package utils

import (
	"path"
	"path/filepath"
	"runtime"
)

const (
	// PortAPITLS for listening api in https
	PortAPITLS string = ":4041"

	// EventTypeNews for event type news
	EventTypeNews string = "NEWS"

	/// Event settings PR

	// EventTitlePR for event title pr
	EventTitlePR string = "Улучшение PR"
	// EventSubTitlePR for event sub title
	EventSubTitlePR string = "Реклама улучшена"
	// EventTypePR for event type pr
	EventTypePR string = "PR"

	/// Event settings Tech

	// EventTitleTech for event title pr
	EventTitleTech string = "Улучшение Tech"
	// EventSubTitleTech for event sub title
	EventSubTitleTech string = "Технологии улучшены"
	// EventTypeTech for event type pr
	EventTypeTech string = "TECH"

	///Payment settings

	// PaymentChips for chips
	PaymentChips string = "payChips"
	// PaymentGolds for golds
	PaymentGolds string = "payGolds"
)

// ProjectPath returns current project path
func ProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) + "/"
}
