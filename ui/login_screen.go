package ui

import (
	"context"
	"edonish-app/client"
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// LoginScreen - экран авторизации
type LoginScreen struct {
	app      fyne.App
	window   fyne.Window
	client   *client.EdonishClient
	username *widget.Entry
	password *widget.Entry
	onLogin  func()
}

// NewLoginScreen создаёт новый экран авторизации
func NewLoginScreen(app fyne.App, win fyne.Window, cli *client.EdonishClient) *LoginScreen {
	return &LoginScreen{
		app:    app,
		window: win,
		client: cli,
	}
}

// SetOnLogin устанавливает коллбек
func (s *LoginScreen) SetOnLogin(callback func()) {
	s.onLogin = callback
}

// Show отображает экран
func (s *LoginScreen) Show() {
	s.username = widget.NewEntry()
	s.username.SetPlaceHolder("Логин")
	s.username.SetFocus(true)

	s.password = widget.NewPasswordEntry()
	s.password.SetPlaceHolder("Пароль")

	loginBtn := widget.NewButton("Войти", func() {
		s.performLogin()
	})

	s.username.OnSubmitted = func(string) {
		s.password.Focus()
	}
	s.password.OnSubmitted = func(string) {
		loginBtn.OnTapped()
	}

	form := container.NewVBox(
		widget.NewLabel("Edonish Pro - Авторизация"),
		widget.NewSeparator(),
		widget.NewForm(
			widget.NewFormItem("Логин", s.username),
			widget.NewFormItem("Пароль", s.password),
		),
		widget.NewSeparator(),
		loginBtn,
	)

	content := container.NewCenter(form)
	s.window.SetContent(content)
	s.window.Resize(fyne.NewSize(400, 300))
	s.window.Show()
}

func (s *LoginScreen) performLogin() {
	username := s.username.Text
	password := s.password.Text

	if username == "" || password == "" {
		dialog.ShowInformation("Ошибка", "Пожалуйста, заполните все поля", s.window)
		return
	}

	disableForm(s.username, s.password)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		resp, err := s.client.Login(ctx, username, password)

		fyne.Do(func() {
			enableForm(s.username, s.password)

			if err != nil {
				var loginErr *client.LoginError
				if errors.As(err, &loginErr) {
					dialog.ShowError("Ошибка авторизации", err, s.window)
				} else {
					dialog.ShowError("Ошибка соединения", err, s.window)
				}
				return
			}

			if !resp.Success {
				dialog.ShowInformation("Ошибка", resp.Message, s.window)
				return
			}

			if s.onLogin != nil {
				s.onLogin()
			}
		})
	}()
}

func disableForm(username, password *widget.Entry) {
	username.Disable()
	password.Disable()
}

func enableForm(username, password *widget.Entry) {
	username.Enable()
	password.Enable()
}
