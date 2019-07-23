package main

import (
	gosxnotifier "github.com/deckarep/gosx-notifier"
)

func pushNotify(msg string) {
	note := gosxnotifier.NewNotification(msg)
	err := note.Push()
	if err != nil {
		logger.Error(err)
	}
}
