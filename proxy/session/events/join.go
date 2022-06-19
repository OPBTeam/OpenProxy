package events

import "github.com/Suremeo/ProxyEye/proxy/session"

type JoinEventHandler func(event *Context, player session.Player)
