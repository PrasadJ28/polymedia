package videostreamer

import "fmt"

type VideoStreamer interface {
	Seek(key string, start, end int) ([]byte, error)
}

type MockVideoStreamer struct {
	Store map[string][]byte
}

func (m *MockVideoStreamer) Seek(key string, start, end int) ([]byte, error) {
	videoData, exist := m.Store[key]
	if !exist {
		return nil, fmt.Errorf("file not found")
	}
	if start < 0 || start >= len(videoData) {
		start = 0
	}
	if end < start || end >= len(videoData) {
		end = len(videoData) - 1
	}
	return videoData[start : end+1], nil
}
