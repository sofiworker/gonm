//go:build ignore
// +build ignore

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"github.com/sofiworker/gonm/libs/notify"
)

func init() {
	notify.Init("notify-example-simple")
}

func main() {
	n := notify.NewNotification("x", "y", "player")
	n.Show()
	n.Destroy()
	notify.Destroy()
}
