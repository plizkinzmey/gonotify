# GoNotify

[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/gonotify.svg)](https://pkg.go.dev/github.com/yourusername/gonotify)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

GoNotify - это простая и легкая библиотека для отправки нативных уведомлений из Go-приложений.

## Особенности

- �� **Нативные уведомления** - использует системные API для отображения уведомлений
- 🍏 **Поддержка macOS** - полная интеграция с macOS Notification Center через UserNotifications API
- 🧩 **Модульный дизайн** - легко расширяемая архитектура
- 🛠️ **Простой API** - минималистичный, понятный интерфейс
- 🧪 **Готов к использованию** - не требует дополнительной настройки

## Требования

- **macOS 10.14+** (Mojave или новее)
- **Go 1.18+**

## Установка

```bash
go get github.com/yourusername/gonotify
```

## Быстрый старт

Создание и отображение простого уведомления:

```go
package main

import "github.com/yourusername/gonotify"

func main() {
    notification := gonotify.New(
        "Заголовок уведомления", 
        "Содержание уведомления",
    )
    
    notification.Show()
}
```

## Подробное использование

### Создание уведомления с иконкой

```go
notification := gonotify.New(
    "Уведомление с иконкой",
    "Это уведомление содержит пользовательскую иконку",
    gonotify.WithIcon("/path/to/icon.png"),
)
notification.Show()
```

### Обработка ошибок

```go
notification := gonotify.New("Заголовок", "Сообщение")
err := notification.Show()
if err != nil {
    log.Fatalf("Ошибка при отображении уведомления: %v", err)
}
```

### Проверка поддержки уведомлений

```go
if gonotify.NotificationsSupported() {
    notification := gonotify.New("Заголовок", "Сообщение")
    notification.Show()
} else {
    log.Println("Нативные уведомления не поддерживаются на этой платформе")
}
```

## Как это работает

### macOS

На macOS библиотека использует UserNotifications framework через CGO и Objective-C для отправки нативных уведомлений.

Уведомления отображаются в Центре уведомлений macOS и могут содержать:
- Заголовок
- Основной текст
- Изображение (опционально)

## Ограничения

- В текущей версии поддерживается только macOS (версии 10.14+)
- Для приложений, запущенных из командной строки, может потребоваться дополнительное разрешение на отправку уведомлений

## Расширение функциональности

Библиотека спроектирована с учётом возможности расширения:

1. Для добавления новых опций уведомления:

```go
// Добавление функции-опции для звука
func WithSound(enable bool) Option {
    return func(n *Notification) {
        n.Sound = enable
    }
}
```

2. Для добавления поддержки других операционных систем, создайте новый файл с соответствующими тегами сборки:

```go
// notify_windows.go
//go:build windows
// +build windows

package gonotify

func showNativeNotification(title, message, iconPath string) error {
    // Реализация для Windows...
}
```

## FAQ

### Как получить доступ к уведомлениям в macOS?

Если ваше приложение запускается из командной строки, вам может понадобиться предоставить ему разрешение на отправку уведомлений в "Системных настройках" -> "Уведомления и фокусировка".

### У меня возникает ошибка при компиляции на не-macOS системе

В текущей версии библиотека поддерживает только macOS. При компиляции на других системах функциональность будет отсутствовать (метод NotificationsSupported() вернет false).

## Интеграция с существующими проектами

### Импортирование модуля

```go
import "github.com/yourusername/gonotify"
```

### Пример использования в Transmission Client

```go
package main

import (
    "github.com/yourusername/gonotify"
    "log"
)

func notifyTorrentCompleted(name string) {
    notification := gonotify.New(
        "Загрузка завершена", 
        "Торрент " + name + " загружен полностью",
    )
    
    if err := notification.Show(); err != nil {
        log.Printf("Ошибка уведомления: %v", err)
    }
}
```

## Лицензия

Распространяется по MIT лицензии. См. файл `LICENSE` для дополнительной информации.
