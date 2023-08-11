package events

import "github.com/opbteam/proxyeye/proxy/session"

type JoinEventHandler func(event *Context, player session.Player)
