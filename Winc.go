package main

import "github.com/tadvi/winc"

func main() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)

	mainWindow.SetText("Jeu du morpion")

	btn := winc.NewPushButton(mainWindow)
	btn0 := winc.NewPushButton(mainWindow)
	btn1 := winc.NewPushButton(mainWindow)
	btn2 := winc.NewPushButton(mainWindow)
	btn3 := winc.NewPushButton(mainWindow)
	btn4 := winc.NewPushButton(mainWindow)
	btn5 := winc.NewPushButton(mainWindow)
	btn6 := winc.NewPushButton(mainWindow)
	btn7 := winc.NewPushButton(mainWindow)
	btn8 := winc.NewPushButton(mainWindow)

	var Jeu string = "Joeur1"

	btn0.SetText(".")
	btn0.SetPos(40, 50)
	btn0.SetSize(100, 40)
	btn0.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn0.SetText("X")
			Jeu = "Joeur2"
		} else if Jeu == "Joueur2" {
			btn0.SetText("O")
			Jeu = "Joeur1"
		}
	})

	btn1.SetText(".")
	btn1.SetPos(40, 100)
	btn1.SetSize(100, 40)
	btn1.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn1.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn1.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn2.SetText(".")
	btn2.SetPos(40, 150)
	btn2.SetSize(100, 40)
	btn2.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn2.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn2.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn3.SetText(".")
	btn3.SetPos(140, 50)
	btn3.SetSize(100, 40)
	btn3.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn3.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn3.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn4.SetText(".")
	btn4.SetPos(140, 100)
	btn4.SetSize(100, 40)
	btn4.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn4.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn4.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn5.SetText(".")
	btn5.SetPos(140, 150)
	btn5.SetSize(100, 40)
	btn5.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn5.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn5.SetText("O")
			Jeu = "Joueur1"
		}
	})

	btn6.SetText(".")
	btn6.SetPos(240, 50)
	btn6.SetSize(100, 40)
	btn6.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn6.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn6.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn7.SetText(".")
	btn7.SetPos(240, 100)
	btn7.SetSize(100, 40)
	btn7.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn7.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn7.SetText("O")
			Jeu = "Joueur1"
		}
	})
	btn8.SetText(".")
	btn8.SetPos(240, 150)
	btn8.SetSize(100, 40)
	btn8.OnClick().Bind(func(e *winc.Event) {
		if Jeu == "Joueur1" {
			btn8.SetText("X")
			Jeu = "Joueur2"
		} else if Jeu == "Joueur2" {
			btn8.SetText("O")
			Jeu = "Joueur1"
		}
	})

	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
}
func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
