package example_test

import (
	"testing"
	"time"

	"github.com/bearchit/goclock/example"

	"github.com/bearchit/goclock"
)

func TestApp_NewUser_Realtime(t *testing.T) {
	clock := goclock.New()
	now := clock.Now()

	app := example.App{
		Clock: clock,
	}

	user := app.NewUser("mock")
	if user.CreatedAt == now {
		t.FailNow()
	}

	t.Logf("now = %v, user.CreatedAt = %v\n", now, user.CreatedAt)
}

func TestApp_NewUser_MockingTime(t *testing.T) {
	clock := goclock.NewMock()
	now := clock.Now()

	app := example.App{
		Clock: clock,
	}

	user := app.NewUser("mock")
	if user.CreatedAt != now {
		t.FailNow()
	}

	t.Logf("now = %v, user.CreatedAt = %v\n", now, user.CreatedAt)
}

func TestApp_NewUser_MockingTime_Controlling(t *testing.T) {
	clock := goclock.NewMock(
		goclock.WithNow(time.Date(2020, 6, 11, 23, 0, 0, 0, time.Local)),
	)
	now := clock.Now()

	app := example.App{
		Clock: clock,
	}

	user := app.NewUser("mock")
	if user.CreatedAt != now {
		t.FailNow()
	}
	t.Logf("now = %v, user.CreatedAt = %v\n", now, user.CreatedAt)

	clock.SetNow(time.Date(2020, 6, 12, 0, 0, 0, 0, time.Local))

	if now == clock.Now() {
		t.FailNow()
	}

	t.Logf("now = %v, clock.Now() = %v\n", now, clock.Now())
}
