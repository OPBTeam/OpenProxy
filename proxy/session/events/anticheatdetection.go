package events

import (
	"github.com/opbteam/proxyeye/proxy/session"
	"github.com/opbteam/proxyeye/proxy/session/anticheat"
)

type AnticheatDetectionEventHandler func(event *Context, player session.Player, detection *anticheat.Detection)
