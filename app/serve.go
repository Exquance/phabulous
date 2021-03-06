package app

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/etcinit/phabulous/app/bot"
	"github.com/jacobstr/confer"
)

// ServeService provides the serve command
type ServeService struct {
	Engine  *EngineService    `inject:""`
	Config  *confer.Config    `inject:""`
	Logger  *logrus.Logger    `inject:""`
	Slacker *bot.SlackService `inject:""`
	App     *Phabulous        `inject:""`
}

// Run starts up the HTTP server
func (s *ServeService) Run(c *cli.Context) {
	// Boot the upper layers of the app.
	s.App.Boot(c)

	s.Logger.Infoln("Starting up the server... (a.k.a. coffee time)")

	engine := s.Engine.New()

	go s.Slacker.BootRTM()

	// Figure out which port to use
	port := ":" + strconv.Itoa(s.Config.GetInt("server.port"))

	engine.Run(port)

	s.Logger.Infoln("✔︎ Done!")
}
