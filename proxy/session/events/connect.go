package events

import "github.com/Suremeo/ProxyEye/proxy/session"

type ConnectEventHandler func(event *Context, player session.Player)
