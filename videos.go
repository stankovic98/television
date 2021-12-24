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
// ---------- Exported Function -----------
func getNumOfVids() int {
	return len(allVideos)
}

func getVideos(from, to int) []Video {
	var wantedVids []Video
	if len(allVideos) < to {
		wantedVids = allVideos[from:]
	} else {
		wantedVids = allVideos[from:to]
	}
	for i, vid := range wantedVids {
		if vid.ThumbnailUrl == "" || vid.Name == "" {
			video := fetchVideoByID(vid.ID)
			wantedVids[i] = video
		}
	}
	return wantedVids
}

// ---------- Internal Function -----------
func fetchVideoByID(id string) Video {
	resp, err := http.Get("https://invidio.xamh.de/api/v1/videos/" + id)
	defer resp.Body.Close()
	if err != nil {
		// handle error
		// change invidius instance
		log.Printf("can't get request for %s: %v\n", id, err)
		return Video{ID: id}
	}
	body, err := io.ReadAll(resp.Body)
	video := getNeededValues(body) // if access denied, return error and try on diferent instance
	return video
}

func initAllVideosWithIDs(videos []Video) []Video {
	for _, id := range allVideoIDs {
		if !isVideoAlreadyFetched(videos, id) {
			videos = append(videos, Video{ID:id})	
		}
	}
	return videos
}

func isVideoAlreadyFetched(videos []Video, id string) bool {
	for _, vid := range videos {
		if vid.ID == id {
			return true
		}
	}
	return false
}

// read from serialized file all videos, if not exists, create file and enter videos from hardcoded file
func getAllSavedVideos() []Video {
	var videos []Video
 	dataFile, err := os.Open("videos.gob")
	defer dataFile.Close()
	// Check if error is "no such file or directory"
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		dataFile, err = getFileFromGithub()
		if err != nil {
			log.Printf("can't get videos.gob file from github: %v\n", err)
			videos = initVideos(0, 12)
			err = saveVideos(videos)
			if err != nil {
				log.Printf("can't save videos: %v\n", err)
			}
			return videos
		}
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

func getFileFromGithub() (*os.File, error) {
	// Create the file
	out, err := os.Create("./videos.gob")
	if err != nil {
		return nil, err
	}
	defer out.Close()
	// Get the data
	resp, err := http.Get("https://raw.githubusercontent.com/stankovic98/television/master/example.csv")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func initVideos(from, to int) []Video {
	var videos []Video
	videoIDs := allVideoIDs[from:to]
	for _, vidID := range videoIDs {
		video := fetchVideoByID(vidID)
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