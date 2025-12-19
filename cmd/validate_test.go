package cmd

import (
	"testing"

	"github.com/moistari/rls"
)

func TestReleaseTypeDetection(t *testing.T) {
	tests := []struct {
		name         string
		folderName   string
		expectedType string
	}{
		{
			name:         "WEB movie release",
			folderName:   "THE.MOVIE.2025.1080P.WEB.H264-GRP",
			expectedType: "movie",
		},
		{
			name:         "BluRay movie release",
			folderName:   "The.Movie.2025.720p.BluRay.x264-GRP",
			expectedType: "movie",
		},
		{
			name:         "TV episode release",
			folderName:   "Show.S01E01.1080p.HDTV.H264-GRP",
			expectedType: "episode",
		},
		{
			name:         "Music live stream release",
			folderName:   "Artist-Live_At_Radio-STREAM-01-01-2025-GRP",
			expectedType: "music",
		},
		{
			name:         "Music album release",
			folderName:   "Artist_-_Album-(CATALOG01)-WEB-2025-GRP",
			expectedType: "music",
		},
		{
			name:         "Application release",
			folderName:   "APP.v1.1.Multilingual.Incl.Keyfilemaker-GRP",
			expectedType: "app",
		},
		{
			name:         "eBook release",
			folderName:   "Publisher.Title.2025.RETAiL.ePub.eBook-GRP",
			expectedType: "book",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			release := rls.ParseString(tt.folderName)
			actualType := release.Type.String()
			if actualType != tt.expectedType {
				t.Errorf("ReleaseTypeDetection(%q) = %q, want %q", tt.folderName, actualType, tt.expectedType)
			}
		})
	}
}
