package example

import (
	"time"

	"github.com/bearchit/goclock"
)

type User struct {
	Name      string
	CreatedAt time.Time
}

type App struct {
	Clock goclock.Clock
}

func (app App) NewUser(name string) User {
	return User{
		Name:      name,
		CreatedAt: app.Clock.Now(),
	}
}
