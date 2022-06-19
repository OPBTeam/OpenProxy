package events

import "github.com/Suremeo/ProxyEye/proxy/session"

type LeaveEventHandler func(event *Context, player session.Player)
