/*
 *
 *           ____                        ______
 *          / __ \_________  _  ____  __/ ____/_  _____
 *         / /_/ / ___/ __ \| |/_/ / / / __/ / / / / _ \
 *        / ____/ /  / /_/ />  </ /_/ / /___/ /_/ /  __/
 *       /_/   /_/   \____/_/|_|\__, /_____/\__, /\___/
 *                                /_/         /_/
 *       ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
 *
 *                     Author: Suremeo (github.com/Suremeo)
 *
 *
 */

package session

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type Session interface {
	Destination
	Source
	Server() Server
	Packet(source Source, destination Destination, pk packet.Packet) error
	Close()
	SetConnected()
	IsConnected() bool
}