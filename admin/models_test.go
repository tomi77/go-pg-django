package admin

import (
	"fmt"
	"testing"
	"time"
)

func ExampleLog_String_addition() {
	var log = Log{
		Id:         1,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
		ActionFlag: ADDITION,
	}

	fmt.Println(log)
	// Output:
	// Added Test.
}

func ExampleLog_String_change() {
	var log = Log{
		Id:            2,
		ActionTime:    time.Now(),
		UserId:        1,
		ObjectRepr:    "Test",
		ActionFlag:    CHANGE,
		ChangeMessage: "update",
	}

	fmt.Println(log)
	// Output:
	// Changed Test - update
}

func ExampleLog_String_deletion() {
	var log = Log{
		Id:         3,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
		ActionFlag: DELETION,
	}

	fmt.Println(log)
	// Output:
	// Deleted Test.
}

func TestLog_String(t *testing.T) {
	var log = Log{
		Id:         4,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
	}

	w := "LogEntry Object"
	g := log.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func ExampleLog_IsAddition() {
	var logs = []Log{
		{
			ActionFlag: ADDITION,
		},
		{
			ActionFlag: CHANGE,
		},
		{
			ActionFlag: DELETION,
		},
		{},
	}

	for _, log := range logs {
		fmt.Println(log.IsAddition())
	}
	// Output:
	// true
	// false
	// false
	// false
}

func ExampleLog_IsChange() {
	var logs = []Log{
		{
			ActionFlag: ADDITION,
		},
		{
			ActionFlag: CHANGE,
		},
		{
			ActionFlag: DELETION,
		},
		{},
	}

	for _, log := range logs {
		fmt.Println(log.IsChange())
	}
	// Output:
	// false
	// true
	// false
	// false
}

func ExampleLog_IsDeletion() {
	var logs = []Log{
		{
			ActionFlag: ADDITION,
		},
		{
			ActionFlag: CHANGE,
		},
		{
			ActionFlag: DELETION,
		},
		{},
	}

	for _, log := range logs {
		fmt.Println(log.IsDeletion())
	}
	// Output:
	// false
	// false
	// true
	// false
}
