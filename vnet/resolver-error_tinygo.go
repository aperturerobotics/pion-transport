// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

package vnet

type hostNotFoundError string

func (err hostNotFoundError) Error() string {
	return "lookup " + string(err) + " on vnet resolver: host not found"
}

func errHostNotFound(hostName string) error {
	return hostNotFoundError(hostName)
}
