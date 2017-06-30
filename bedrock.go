package main

import (
    "fmt"
//    "github.com/mattn/go-gtk/gdkpixbuf"
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
    "os"
//    "os/exec"
//    "path/filepath"
//    "regexp"
//    "sort"
//    "strings"
)


func main() {
    fmt.Println("Hello World")

    gtk.Init(&os.Args)

    window := gtk.NewWindow(gtk.WINDOW_POPUP)
    window.SetName("Bedrock")
    window.SetDecorated(false)
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        fmt.Println("got destroy!", ctx.Data().(string))
        gtk.MainQuit()
    }, "foo")

    window.SetSizeRequest(600,40)
    window.ShowAll()
    gtk.Main()
}
