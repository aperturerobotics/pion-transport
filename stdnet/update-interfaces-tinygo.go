// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

package stdnet

func (n *Net) updateInterfaces() error {
	n.interfaces = nil
	return nil
}
