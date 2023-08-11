package events

import "github.com/opbteam/proxyeye/proxy/session"

type LeaveEventHandler func(event *Context, player session.Player)
