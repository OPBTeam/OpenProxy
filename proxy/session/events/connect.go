package events

import "github.com/opbteam/proxyeye/proxy/session"

type ConnectEventHandler func(event *Context, player session.Player)
