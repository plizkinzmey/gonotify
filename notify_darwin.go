//go:build darwin
// +build darwin

package gonotify

/*
#cgo CFLAGS: -mmacosx-version-min=10.14 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.14 -framework Foundation -framework UserNotifications
#include <stdlib.h>
#include "darwin/NotificationBridge.h"
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
