// Package gonotify предоставляет кроссплатформенный интерфейс
// для отправки нативных уведомлений из Go-приложений
package gonotify

// Notification представляет нативное уведомление
type Notification struct {
	Title    string // Заголовок уведомления
	Message  string // Содержание уведомления
	IconPath string // Путь к иконке (если поддерживается платформой)
}

// Option - функция для настройки уведомления
type Option func(*Notification)

// New создает новое уведомление с заданными заголовком и сообщением
func New(title, message string, opts ...Option) *Notification {
	n := &Notification{
		Title:   title,
		Message: message,
	}

	// Применяем все переданные опции
	for _, opt := range opts {
		opt(n)
	}

	return n
}

// WithIcon добавляет иконку к уведомлению
func WithIcon(iconPath string) Option {
	return func(n *Notification) {
		n.IconPath = iconPath
	}
}

// Show отображает уведомление с использованием нативного API
// текущей операционной системы
func (n *Notification) Show() error {
	return showNativeNotification(n.Title, n.Message, n.IconPath)
}
