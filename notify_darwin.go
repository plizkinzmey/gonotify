//go:build darwin
// +build darwin

package gonotify

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.14 -I${SRCDIR}/darwin
#cgo LDFLAGS: -framework Foundation -framework UserNotifications -mmacosx-version-min=10.14

#include <stdlib.h>
#include "NotificationBridge.h"
#include "NotificationBridge.m"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Реализация showNativeNotification для macOS
func showNativeNotification(title, message, iconPath string) error {
	ctitle := C.CString(title)
	cmessage := C.CString(message)
	cicon := C.CString(iconPath)
	defer C.free(unsafe.Pointer(ctitle))
	defer C.free(unsafe.Pointer(cmessage))
	defer C.free(unsafe.Pointer(cicon))

	result := C.ShowNativeNotification(ctitle, cmessage, cicon)
	if result != 0 {
		return errors.New("failed to show notification")
	}

	return nil
}

// NotificationsSupported возвращает true, так как macOS поддерживает уведомления
func NotificationsSupported() bool {
	return true
}
