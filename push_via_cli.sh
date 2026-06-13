#!/bin/bash

# Отправка через GitHub CLI
# Установите gh если не установлен: https://cli.github.com/

if ! command -v gh &> /dev/null; then
    echo "GitHub CLI не установлен"
    echo "Установите: https://cli.github.com/"
    echo ""
    echo "Ubuntu/Debian:"
    echo "  sudo apt install gh"
    echo ""
    echo "macOS:"
    echo "  brew install gh"
    echo ""
    echo "После установки:"
    echo "  gh auth login"
    echo "  ./push_via_cli.sh"
    exit 1
fi

# Авторизация если не авторизованы
if ! gh auth status &> /dev/null; then
    echo "🔐 Авторизация на GitHub..."
    gh auth login
fi

echo "🚀 Отправка через GitHub CLI..."
gh repo create eDonish_pro --public --source=. --push --confirm

if [ $? -eq 0 ]; then
    echo "✅ Успешно!"
else
    echo "❌ Ошибка"
fi
