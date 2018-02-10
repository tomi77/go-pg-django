package admin

import (
	"fmt"
	"testing"
	"time"
)

var logs = []Log{
	{
		Id:         1,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
		ActionFlag: ADDITION,
	},
	{
		Id:            2,
		ActionTime:    time.Now(),
		UserId:        1,
		ObjectRepr:    "Test",
		ActionFlag:    CHANGE,
		ChangeMessage: "update",
	},
	{
		Id:         3,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
		ActionFlag: DELETION,
	},
	{
		Id:         4,
		ActionTime: time.Now(),
		UserId:     1,
		ObjectRepr: "Test",
	},
}

func TestString(t *testing.T) {
	t.Run("ADDITION", func(t *testing.T) {
		g := logs[0].String()
		w := "Added Test."
		if g != w {
			t.Error(fmt.Sprintf("Wanted %q got %q", w, g))
		}
	})
	t.Run("CHANGE", func(t *testing.T) {
		g := logs[1].String()
		w := "Changed Test - update"
		if g != w {
			t.Error(fmt.Sprintf("Wanted %q got %q", w, g))
		}
	})
	t.Run("DELETION", func(t *testing.T) {
		g := logs[2].String()
		w := "Deleted Test."
		if g != w {
			t.Error(fmt.Sprintf("Wanted %q got %q", w, g))
		}
	})
	t.Run("other", func(t *testing.T) {
		g := logs[3].String()
		w := "LogEntry Object"
		if g != w {
			t.Error(fmt.Sprintf("Wanted %q got %q", w, g))
		}
	})
}

func TestIsAddition(t *testing.T) {
	t.Run("ADDITION", func(t *testing.T) {
		if logs[0].IsAddition() != true {
			t.Error("Wanted true got false")
		}
	})
	t.Run("CHANGE", func(t *testing.T) {
		if logs[1].IsAddition() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("DELETION", func(t *testing.T) {
		if logs[2].IsAddition() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("other", func(t *testing.T) {
		if logs[3].IsAddition() != false {
			t.Error("Wanted false got true")
		}
	})
}

func TestIsChange(t *testing.T) {
	t.Run("ADDITION", func(t *testing.T) {
		if logs[0].IsChange() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("CHANGE", func(t *testing.T) {
		if logs[1].IsChange() != true {
			t.Error("Wanted true got false")
		}
	})
	t.Run("DELETION", func(t *testing.T) {
		if logs[2].IsChange() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("other", func(t *testing.T) {
		if logs[3].IsChange() != false {
			t.Error("Wanted false got true")
		}
	})
}

func TestIsDeletion(t *testing.T) {
	t.Run("ADDITION", func(t *testing.T) {
		if logs[0].IsDeletion() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("CHANGE", func(t *testing.T) {
		if logs[1].IsDeletion() != false {
			t.Error("Wanted false got true")
		}
	})
	t.Run("DELETION", func(t *testing.T) {
		if logs[2].IsDeletion() != true {
			t.Error("Wanted true got false")
		}
	})
	t.Run("other", func(t *testing.T) {
		if logs[3].IsDeletion() != false {
			t.Error("Wanted false got true")
		}
	})
}
