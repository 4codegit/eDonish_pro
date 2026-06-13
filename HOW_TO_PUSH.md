# 📤 Как загрузить код на GitHub

## Способ 1: Через веб-интерфейс (самый простой)

1. Перейдите на: **https://github.com/4codegit/eDonish_pro**

2. Нажмите **"uploading an existing file"**

3. Перетащите все файлы из директории `edonish_pro/`:
   - `.gitignore`
   - `README.md`
   - `go.mod`
   - `main.go`
   - `controller.go`
   - `client/client.go`
   - `ui/login_screen.go`
   - `ui/dashboard.go`

4. Добавьте комментарий: `Initial commit`

5. Нажмите **"Commit changes"**

---

## Способ 2: Через Git (если настроен SSH)

```bash
cd /home/narziev/dev/edonish-auto/edonish_pro

# Если SSH ключ настроен:
git remote add origin git@github.com:4codegit/eDonish_pro.git
git push -u origin main
```

---

## Способ 3: Через GitHub CLI

```bash
# Установите GitHub CLI
sudo apt install gh

# Авторизуйтесь
gh auth login

# Отправьте код
gh repo create eDonish_pro --public --source=. --push --confirm
```

---

## Проверка

После загрузки проверьте: **https://github.com/4codegit/eDonish_pro**

Все файлы должны быть видны в репозитории.
