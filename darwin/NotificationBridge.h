#import <Foundation/Foundation.h>

#ifdef __cplusplus
extern "C" {
#endif

// ShowNativeNotification отображает нативное уведомление в macOS
// Возвращает 0 при успехе, отличное от 0 значение при ошибке
int ShowNativeNotification(const char *title, const char *message, const char *iconPath);

#ifdef __cplusplus
}
#endif
