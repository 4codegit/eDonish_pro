# eDonish Pro 📚

**Профессиональная версия автоматизации электронного журнала edonish.tj**

Десктопное приложение на **Go** + **Fyne GUI**

---

## 🚀 Скачать

### GitHub Releases

[Скачать последнюю версию](https://github.com/4codegit/eDonish_pro/releases)

| Платформа | Файл |
|-----------|------|
| 🐧 Linux (DEB) | `edonish-pro_x.x.x_amd64.deb` |
| 🐧 Linux (RPM) | `edonish-pro-x.x.x-1.x86_64.rpm` |
| 🪟 Windows | `edonish-pro-windows.exe` |
| 🪟 Windows Setup | `edonish-pro-windows-installer.exe` |
| 🤖 Android | `edonish-pro.apk` |

---

## 🏃 Быстрый старт

### Установка зависимостей

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install -y gcc libgl1-mesa-dev xorg-dev
```

**macOS:**
```bash
brew install go
brew install gcc
```

**Windows:**
```bash
choco install mingw
```

### Запуск

```bash
# Клонируйте репозиторий
git clone https://github.com/4codegit/eDonish_pro.git
cd eDonish_pro

# Установите зависимости
go mod tidy

# Запустите
go run .
```

### Компиляция

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o edonish-pro .

# Windows
GOOS=windows GOARCH=amd64 go build -o edonish-pro.exe .

# macOS
GOOS=darwin GOARCH=amd64 go build -o edonish-pro .
```

---

## 📋 Структура

```
eDonish_pro/
├── main.go              # Точка входа
├── controller.go        # Контроллер
├── client/
│   └── client.go        # API клиент
├── ui/
│   ├── login_screen.go  # Вход
│   └── dashboard.go     # Дашборд
├── go.mod
└── README.md
```

---

## ⚙️ Настройка API

Откройте `controller.go` и настройте endpoints:

```go
c.client, err = client.NewEdonishClient(
    "https://edonish.tj",
    "https://edonish.tj/auth/v1/login", // ← Замените на реальный
)
```

---

## 📦 CI/CD

Автоматическая сборка при теге:

```bash
git tag v0.1.0
git push origin v0.1.0
```

Файлы появятся в [Releases](https://github.com/4codegit/eDonish_pro/releases)

---

## 🔐 Учётные данные

```
Логин: 200117707
Пароль: ********
```

**Внимание:** Никогда не коммитьте реальные учётные данные!

---

## 📝 Лицензия

MIT License - образовательные цели.
