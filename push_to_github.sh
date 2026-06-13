#!/bin/bash

# Скрипт для отправки кода на GitHub
# Запустите после создания репозитория

echo "🚀 Отправка кода на GitHub..."

# Удалите существующий remote если есть
git remote remove origin 2>/dev/null

# Добавьте новый remote
git remote add origin git@github.com:4codegit/eDonish_pro.git

# Отправьте код
git push -u origin main

if [ $? -eq 0 ]; then
    echo "✅ Успешно отправлено!"
    echo "📱 Проверьте: https://github.com/4codegit/eDonish_pro"
else
    echo "❌ Ошибка при отправке"
    echo ""
    echo "Попробуйте через GitHub CLI:"
    echo "  gh repo create eDonish_pro --public --source=. --push"
    echo ""
    echo "Или через веб-интерфейс:"
    echo "  1. Перейдите на https://github.com/4codegit/eDonish_pro"
    echo "  2. Нажмите 'uploading an existing file'"
    echo "  3. Перетащите все файлы из этой директории"
fi
