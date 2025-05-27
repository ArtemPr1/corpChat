package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

const (
	AppName    = "CorpChat"
	AppVersion = "0.0.1"
)

func main() {
	// Инициализация приложения
	application := app.New()

	// Создание главного окна
	window := application.NewWindow(AppName + " v" + AppVersion)
	window.SetMaster() // Главное окно (закрытие = выход из приложения)

	// Загрузка конфигов (если есть)
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация сервисов (БД, API и т.д.)
	services, err := initServices(cfg)
	if err != nil {
		log.Fatalf("Failed to init services: %v", err)
	}
	defer services.Close() // Закрытие соединений при выходе

	// Создание UI
	ui := createUI(services, window)
	window.SetContent(ui)

	// Старт приложения
	window.ShowAndRun()
}

// --- Конфигурация ---
type Config struct {
	APIToken  string
	ServerURL string
}

func loadConfig() (*Config, error) {
	// Пример: загрузка из .env или файла конфига
	return &Config{
		APIToken:  "your_api_token",
		ServerURL: "http://localhost:8080",
	}, nil
}

// --- Сервисы ---
type Services struct {
	// Здесь могут быть: клиент API, БД и т.д.
}

func initServices(cfg *Config) (*Services, error) {
	// Пример: инициализация подключения к API/БД
	return &Services{}, nil
}

func (s *Services) Close() {
	// Закрытие всех соединений
	log.Println("Closing services...")
}

// --- UI ---
func createUI(services *Services, window fyne.Window) fyne.CanvasObject {
	// Пример: создание интерфейса (логин/чат)
	loginBtn := widget.NewButton("Login", func() {
		log.Println("Login clicked!")
		showChatWindow(services, window)
	})

	return container.NewVBox(
		widget.NewLabel("Welcome to CorpChat!"),
		loginBtn,
	)
}

func showChatWindow(services *Services, window fyne.Window) {
	// Пример: окно чата
	chatInput := widget.NewEntry()
	chatInput.SetPlaceHolder("Type a message...")

	messages := widget.NewLabel("Messages will appear here...")

	sendBtn := widget.NewButton("Send", func() {
		msg := chatInput.Text
		if msg == "" {
			return
		}
		log.Printf("Sending message: %s", msg)
		// Здесь будет вызов API для отправки сообщения
		messages.SetText(messages.Text + "\nYou: " + msg)
		chatInput.SetText("")
	})

	chatContent := container.NewVScroll(messages)
	chatContent.SetMinSize(fyne.NewSize(300, 200))

	window.SetContent(container.NewVBox(
		chatContent,
		container.NewBorder(nil, nil, nil, sendBtn, chatInput),
	))
}
