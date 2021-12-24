package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

func basic() string {
  return "World!"
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

var allVideoIDs = getVideoIDs()

// read all videos from .gob file, if video with ID is not present in file just fill out
// videoIDs, and when request is made for video that has only ID, then call API for data
var allVideos = getAllSavedVideos()

func main() {
  allVideos = initAllVideosWithIDs(allVideos)
  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "Television",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Bind(getVideos)
  app.Bind(getNumOfVids)
  app.Run()
}
