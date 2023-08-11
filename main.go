package main

import (
	"os"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/opbteam/proxyeye/proxy"
	"github.com/opbteam/proxyeye/proxy/console"
	"github.com/opbteam/proxyeye/proxy/session"
	"github.com/opbteam/proxyeye/proxy/session/anticheat"
	"github.com/opbteam/proxyeye/proxy/session/events"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func main() {
	proxy.Eye.HandleEvent(events.JoinEventHandler(func(event *events.Context, player session.Player) {
		console.Info("Player connected: %v", player.Raknet().IdentityData().DisplayName)
	}))
	proxy.Eye.HandleEvent(events.LeaveEventHandler(func(event *events.Context, player session.Player) {
		console.Warn("Player left: %v", player.Raknet().IdentityData().DisplayName)
	}))
	one := proxy.NewRemoteServer("", 100)
	two := proxy.NewInternalServer()
	proxy.Eye.HandleEvent(events.ConnectEventHandler(func(event *events.Context, player session.Player) {
		console.Debug("%v connected to %v", player.Raknet().IdentityData().DisplayName, player.Session().Server().Address())
		time.Sleep(30 * time.Second)
		if player.Session().Server().Address() == one.Address() {
			_ = two.Connect(player)
		} else {
			_ = one.Connect(player)
		}
	}))
	proxy.Eye.HandleEvent(events.AnticheatDetectionEventHandler(func(event *events.Context, player session.Player, detection *anticheat.Detection) {
		console.Debug("%v is suspected of using %v", player.Raknet().IdentityData().DisplayName, detection.Type)
	}))
	proxy.Eye.HandleEvent(events.PacketEventHandler(func(event *events.Context, player session.Player, source session.Source, pk packet.Packet) {

	}))
	go func() {
		i := 0
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			i++
			c, err := os.Create("./heaps/" + strconv.Itoa(i) + "-heap.pprof")
			if err != nil {
				console.Error("Failed to create heap file", err)
			} else {
				err = pprof.WriteHeapProfile(c)
				if err != nil {
					console.Error("Failed to write to heap file", err)
				} else {
					console.Info("Wrote heap %v", i)
				}
			}
		}
	}()
	console.Fatal("Error starting proxy", proxy.Eye.Run())
}
