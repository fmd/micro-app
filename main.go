package main

import (
    "github.com/go-martini/martini"
)

type App struct {
    Martini *martini.ClassicMartini
}

func (a *App) Init() {
    a.Martini = martini.Classic()
    a.Martini.Get("/", func() (int, string) {
        return 200, "OK!"
    })
    a.Martini.Get("/error", func() (int, string) {
        return 500, "OH NO!"
    })
    a.Martini.Post("/submit", func() (int, string) {
        return 401, "NUH HUH!"
    })
}

func (a *App) Run() {
    a.Martini.Run()
}

func main() {
    a := &App{}
    a.Init()
    a.Run()
}
