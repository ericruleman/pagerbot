package updater

import (
	"github.com/karlkfi/pagerbot/internal/config"
	"github.com/karlkfi/pagerbot/internal/pagerduty"
	"github.com/karlkfi/pagerbot/internal/slack"
	"sync"
	"time"
)

type Updater struct {
	Wg        *sync.WaitGroup
	Slack     *slack.Api
	Pagerduty *pagerduty.Api
	Users     *UserList
	Schedules *ScheduleList
	LastFetch time.Time
}

func New() (*Updater, error) {
	u := Updater{}
	u.Wg = &sync.WaitGroup{}

	var err error
	u.Slack, err = slack.New(config.Config.ApiKeys.Slack)
	if err != nil {
		return &u, err
	}

	u.Pagerduty, err = pagerduty.New(config.Config.ApiKeys.Pagerduty.Key, config.Config.ApiKeys.Pagerduty.Org)
	if err != nil {
		return &u, err
	}

	u.Users = &UserList{}
	u.Schedules = &ScheduleList{}

	return &u, nil
}
