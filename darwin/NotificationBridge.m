// NotificationBridge.m - Интерфейс для показа нативных уведомлений в macOS
#import "NotificationBridge.h"
#import <Foundation/Foundation.h>
#import <UserNotifications/UserNotifications.h>

// Для UserNotifications нужен минимум macOS 10.14
int ShowNativeNotification(const char *title, const char *message, const char *iconPath) {
    __block int success = 0;
    @autoreleasepool {
        NSLog(@"[GoNotify] ShowNativeNotification: %s | %s", title, message);
        NSString *nsTitle = [NSString stringWithUTF8String:title];
        NSString *nsMessage = [NSString stringWithUTF8String:message];
        
        UNMutableNotificationContent *content = [[UNMutableNotificationContent alloc] init];
        content.title = nsTitle;
        content.body = nsMessage;
        
        // Добавляем иконку, если указан путь
        if (iconPath != NULL && strlen(iconPath) > 0) {
            NSString *nsIconPath = [NSString stringWithUTF8String:iconPath];
            if ([[NSFileManager defaultManager] fileExistsAtPath:nsIconPath]) {
                NSURL *iconURL = [NSURL fileURLWithPath:nsIconPath];
                NSError *error = nil;
                UNNotificationAttachment *attachment = [UNNotificationAttachment attachmentWithIdentifier:@"icon"
                                                                                                     URL:iconURL
                                                                                                 options:nil
                                                                                                   error:&error];
                if (error) {
                    NSLog(@"[GoNotify] Ошибка добавления иконки: %@", error);
                } else if (attachment) {
                    content.attachments = @[attachment];
                }
            }
        }
        
        UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
        dispatch_semaphore_t sem = dispatch_semaphore_create(0);
        
        [center requestAuthorizationWithOptions:(UNAuthorizationOptionAlert | UNAuthorizationOptionSound | UNAuthorizationOptionBadge)
                              completionHandler:^(BOOL granted, NSError * _Nullable error) {
            NSLog(@"[GoNotify] Разрешения: granted=%d, error=%@", granted, error);
            
            if (!granted) {
                NSLog(@"[GoNotify] Нет разрешения на уведомления");
                success = 1; // Ошибка - нет разрешения
                dispatch_semaphore_signal(sem);
                return;
            }
            
            UNNotificationRequest *request = [UNNotificationRequest requestWithIdentifier:[[NSUUID UUID] UUIDString] 
                                                                                  content:content 
                                                                                  trigger:nil];
            
            [center addNotificationRequest:request withCompletionHandler:^(NSError * _Nullable error) {
                if (error) {
                    NSLog(@"[GoNotify] Ошибка при добавлении уведомления: %@", error);
                    success = 2; // Ошибка при добавлении уведомления
                } else {
                    NSLog(@"[GoNotify] Уведомление успешно добавлено");
                }
                dispatch_semaphore_signal(sem);
            }];
        }];
        
        // Ожидаем завершения операций с таймаутом
        dispatch_semaphore_wait(sem, dispatch_time(DISPATCH_TIME_NOW, (int64_t)(3 * NSEC_PER_SEC)));
    }
    
    return success;
}
