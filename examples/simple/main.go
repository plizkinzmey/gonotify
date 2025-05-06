package main

import (
"fmt"
"os"
"path/filepath"
"time"

"github.com/yourusername/gonotify" // В реальном проекте здесь будет ваш репозиторий
)

func main() {
	fmt.Println("GoNotify - Пример использования")
	
	// Создание простого уведомления
	notification := gonotify.New(
"Пример уведомления",
"Это пример использования модуля gonotify",
)
	
	// Отправка уведомления
	err := notification.Show()
	if err != nil {
		fmt.Printf("Ошибка при отображении уведомления: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("Уведомление отправлено!")
	
	// Подождем немного
	time.Sleep(1 * time.Second)
	
	// Пример с иконкой
	execPath, err := os.Executable()
	if err == nil {
		iconPath := filepath.Join(filepath.Dir(execPath), "icon.png")
		
		// Создание уведомления с иконкой
		notificationWithIcon := gonotify.New(
"Уведомление с иконкой",
"Это уведомление содержит пользовательскую иконку",
gonotify.WithIcon(iconPath), // Использование опции для добавления иконки
)
		
		// Отправка уведомления
		err = notificationWithIcon.Show()
		if err != nil {
			fmt.Printf("Ошибка при отображении уведомления с иконкой: %v\n", err)
		} else {
			fmt.Println("Уведомление с иконкой отправлено!")
		}
	}
}
