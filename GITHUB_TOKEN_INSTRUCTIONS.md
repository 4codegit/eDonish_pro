# 📝 Инструкция по созданию GitHub Token

## Шаг 1: Создайте токен

1. Перейдите на: https://github.com/settings/tokens
2. Нажмите **"Generate new token (classic)"**
3. Дайте название: `eDonish_pro CI/CD`
4. Выберите срок: `No expiration`
5. Выберите scope:
   - ✅ `repo` (Full control of private repositories)
   - ✅ `workflow` (Update GitHub Action workflow files)
6. Нажмите **"Generate token"**
7. **Скопируйте токен** (он показывается только один раз!)

## Шаг 2: Используйте токен

```bash
# Установите переменную окружения
export GITHUB_TOKEN="ваш_токен_сюда"

# Или используйте в команде:
git remote set-url origin https://TOKEN@github.com/4codegit/eDonish_pro.git
git push -u origin main
```

## Шаг 3: Альтернатива - GitHub CLI

```bash
# Установите GitHub CLI (если есть sudo)
sudo apt install gh

# Авторизуйтесь
gh auth login --with-token <<< "ваш_токен"

# Отправьте код
cd /home/narziev/dev/edonish-auto/edonish_pro
gh repo create eDonish_pro --public --source=. --push --confirm
```

---

**ВАЖНО:** Никогда не делитесь токеном и не коммитьте его в репозиторий!
