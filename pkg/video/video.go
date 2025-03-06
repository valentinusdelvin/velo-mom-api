package video

import "strings"

func ExtractVideoID(url string) string {
	if strings.Contains(url, "youtu.be/") {
		parts := strings.Split(url, "youtu.be/")
		if len(parts) > 1 {
			return strings.Split(parts[1], "?")[0]
		}
	} else if strings.Contains(url, "youtube.com/watch?v=") {
		parts := strings.Split(url, "v=")
		if len(parts) > 1 {
			return strings.Split(parts[1], "&")[0]
		}
	}
	return ""
}

func GenerateThumbnail(youtubeID string) string {
	if youtubeID == "" {
		return ""
	}
	return "https://img.youtube.com/vi/" + youtubeID + "/hqdefault.jpg"
}
