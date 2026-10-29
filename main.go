package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"tome/character"
	"tome/store"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// dbPath returns the OS equivalent of Go config location (~/.config/tome/tome.db for linux)
func dbPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find config dir: %w", err)
	}

	dir = filepath.Join(dir, "tome")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create %s: %w", dir, err)
	}

	return filepath.Join(dir, "tome.db"), nil
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// character setup
	path, err := dbPath()
	if err != nil {
		log.Fatal(err)
	}
	st, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
	characterSvc := character.NewCharacterService(st)
	characterApi := NewCharacterApi(characterSvc)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "tome",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
			characterApi,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
