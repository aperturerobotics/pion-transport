// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package stdnet

import (
	"github.com/pion/transport/v4"
	"github.com/wlynxg/anet"
)

func (n *Net) updateInterfaces() error {
	ifs := []*transport.Interface{}

	// anet works around Android 11+ netlink restrictions affecting
	// net.Interfaces and net.InterfaceAddrs. See golang/go#40569.
	oifs, err := anet.Interfaces()
	if err != nil {
		return err
	}

	for i := range oifs {
		ifc := transport.NewInterface(oifs[i])

		addrs, err := anet.InterfaceAddrsByInterface(&oifs[i])
		if err != nil {
			return err
		}

		for _, addr := range addrs {
			ifc.AddAddress(addr)
		}

		ifs = append(ifs, ifc)
	}

	n.interfaces = ifs

	return nil
}
