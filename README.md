# eDonish Pro

Автоматизированная система для **edonish.tj**

## Статус

✅ CI/CD настроен - автоматическая сборка при создании тега

## Установка

### Linux
```bash
# Скачать из GitHub Releases
wget https://github.com/4codegit/eDonish_pro/releases/latest/download/edonish-pro
chmod +x edonish-pro
```

## Использование

### Консольная версия
```bash
./edonish-pro <username> <password>
```

Пример:
```bash
./edonish-pro 200117707 test123
```

### GUI версия (требует Fyne)
```bash
go run -tags gui main_gui.go
```

## Сборка

### Локальная
```bash
go build -o edonish-pro .
```

### CI/CD
Автоматически при создании тега:
```bash
git tag v0.2.0
git push origin v0.2.0
```

## Структура
- `main.go` - основное приложение (консоль)
- `main_gui.go` - GUI версия (Fyne, опционально)
- `.github/workflows/release.yml` - CI/CD конфигурация

## Технологии
- Go 1.21
- Fyne v2 (GUI, опционально)
- GitHub Actions (CI/CD)

## License
MIT
