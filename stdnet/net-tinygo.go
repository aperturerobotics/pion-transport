// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

// Package stdnet implements the transport.Net interface
// using methods from Go's standard net package.
package stdnet

import (
	"context"
	"fmt"
	"net"

	"github.com/pion/transport/v4"
)

const (
	lo0String = "lo0String"
	udpString = "udp"
)

// Net is an implementation of the net.Net interface
// based on functions of the standard net package.
type Net struct {
	interfaces []*transport.Interface
}

// NewNet creates a new StdNet instance.
func NewNet() (*Net, error) {
	n := &Net{}

	return n, n.UpdateInterfaces()
}

// UpdateInterfaces updates the internal list of network interfaces
// and associated addresses.
func (n *Net) UpdateInterfaces() error {
	return n.updateInterfaces()
}

// Interfaces returns a slice of interfaces which are available on the
// system.
func (n *Net) Interfaces() ([]*transport.Interface, error) {
	return n.interfaces, nil
}

// InterfaceByIndex returns the interface specified by index.
func (n *Net) InterfaceByIndex(index int) (*transport.Interface, error) {
	for _, ifc := range n.interfaces {
		if ifc.Index == index {
			return ifc, nil
		}
	}

	return nil, fmt.Errorf("%w: index=%d", transport.ErrInterfaceNotFound, index)
}

// InterfaceByName returns the interface specified by name.
func (n *Net) InterfaceByName(name string) (*transport.Interface, error) {
	for _, ifc := range n.interfaces {
		if ifc.Name == name {
			return ifc, nil
		}
	}

	return nil, fmt.Errorf("%w: %s", transport.ErrInterfaceNotFound, name)
}

// ListenPacket announces on the local network address.
func (n *Net) ListenPacket(network string, address string) (net.PacketConn, error) {
	return nil, transport.ErrNotSupported
}

// ListenUDP acts like ListenPacket for UDP networks.
func (n *Net) ListenUDP(network string, locAddr *net.UDPAddr) (transport.UDPConn, error) {
	return nil, transport.ErrNotSupported
}

// Dial connects to the address on the named network.
func (n *Net) Dial(network, address string) (net.Conn, error) {
	return nil, transport.ErrNotSupported
}

// DialUDP acts like Dial for UDP networks.
func (n *Net) DialUDP(network string, laddr, raddr *net.UDPAddr) (transport.UDPConn, error) {
	return nil, transport.ErrNotSupported
}

// ResolveIPAddr returns an address of IP end point.
func (n *Net) ResolveIPAddr(network, address string) (*net.IPAddr, error) {
	return net.ResolveIPAddr(network, address)
}

// ResolveUDPAddr returns an address of UDP end point.
func (n *Net) ResolveUDPAddr(network, address string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr(network, address)
}

// ResolveTCPAddr returns an address of TCP end point.
func (n *Net) ResolveTCPAddr(network, address string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr(network, address)
}

// DialTCP acts like Dial for TCP networks.
func (n *Net) DialTCP(network string, laddr, raddr *net.TCPAddr) (transport.TCPConn, error) {
	return nil, transport.ErrNotSupported
}

// ListenTCP acts like Listen for TCP networks.
func (n *Net) ListenTCP(network string, laddr *net.TCPAddr) (transport.TCPListener, error) {
	return nil, transport.ErrNotSupported
}

type stdDialer struct{}

func (d stdDialer) Dial(network, address string) (net.Conn, error) {
	return nil, transport.ErrNotSupported
}

// CreateDialer creates an instance of vnet.Dialer.
func (n *Net) CreateDialer(d *net.Dialer) transport.Dialer {
	return stdDialer{}
}

type stdListenConfig struct{}

func (d stdListenConfig) Listen(ctx context.Context, network, address string) (net.Listener, error) {
	return nil, transport.ErrNotSupported
}

func (d stdListenConfig) ListenPacket(ctx context.Context, network, address string) (net.PacketConn, error) {
	return nil, transport.ErrNotSupported
}

// CreateListenConfig creates an instance of vnet.ListenConfig.
func (n *Net) CreateListenConfig(d *net.ListenConfig) transport.ListenConfig {
	return stdListenConfig{}
}

// Compile-time assertion.
var _ transport.Net = &Net{}
