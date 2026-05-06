// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package vnet

import "net"

func errHostNotFound(hostName string) error {
	return &net.DNSError{
		Err:         "host not found",
		Name:        hostName,
		Server:      "vnet resolver",
		IsTimeout:   false,
		IsTemporary: false,
	}
}
