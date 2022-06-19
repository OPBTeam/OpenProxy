package events

import (
	"github.com/Suremeo/ProxyEye/proxy/session"
	"github.com/Suremeo/ProxyEye/proxy/session/anticheat"
)

type AnticheatDetectionEventHandler func(event *Context, player session.Player, detection *anticheat.Detection)
