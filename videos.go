package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
)

type Video struct {
	ID	string
	Date string
	Name string
	ThumbnailUrl string
}

// read from serialized file all videos, if not exists, create file and enter videos from hardcoded file
func GetAllVideos() []Video {
	var videos []Video
 	dataFile, err := os.Open("videos.gob")
	defer dataFile.Close()
	// Check if error is "no such file or directory"
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		videos = initVideos()
		err := saveVideos(videos)
		if err != nil {
			log.Printf("can't save videos: %v\n", err)
		}
		return videos
	}
 	if err != nil {
 		fmt.Println(err)
 		return nil
 	}
 	dataDecoder := gob.NewDecoder(dataFile)
 	err = dataDecoder.Decode(&videos)
 	if err != nil {
 		fmt.Println(err)
 		return nil
 	}
	return videos
}

func saveVideos(videos []Video) error {
	dataFile, err := os.Create("videos.gob")
 	if err != nil {
 		fmt.Println(err)
		return err
	}
 	dataEncoder := gob.NewEncoder(dataFile)
 	dataEncoder.Encode(videos)
 	dataFile.Close()
	return nil
}

func getHardcodedVids() []Video {
	return nil
}

func initVideos() []Video {
	var videos []Video
	videoIDs := getVideoIDs()
	for _, vidID := range videoIDs {
		resp, err := http.Get("https://invidio.xamh.de/api/v1/videos/" + vidID)
		defer resp.Body.Close()
		if err != nil {
			// handle error
			// change invidius instance
			log.Printf("can't get request for %s: %v\n", vidID, err)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		video := getNeededValues(body) // if access denied, return error and try on diferent instance
		videos = append(videos, video)
	}
	return videos
}

type InvVideoStruct struct {
	Title string `json:"title"`
	ID string `json:"videoId"`
	VidThumbnails []struct{
		Url string `json:"url"`
	} `json:"videoThumbnails"`  
} 

func getNeededValues(body []byte) Video {
	var invVidS InvVideoStruct 
	err := json.Unmarshal(body, &invVidS)
	if err != nil {
		log.Printf("can't unmarshal response body: %v\n", err)
		return Video{}
	}
	log.Println("it works", invVidS.VidThumbnails[0].Url)
	return Video{Name: invVidS.Title, ID: invVidS.ID, ThumbnailUrl: invVidS.VidThumbnails[4].Url}
}

func getVideoIDs() []string {
	var videoIDs []string
	file, err := os.Open("hardcoded-videos")
    if err != nil {
        log.Println(err)
		return nil
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		if len(line) == 11 {
			videoIDs = append(videoIDs, line)
		} else {
			vidID := strings.SplitAfter(line, "/watch?v=")[1]
			videoIDs = append(videoIDs, vidID[:11])
		}
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
	return videoIDs
}