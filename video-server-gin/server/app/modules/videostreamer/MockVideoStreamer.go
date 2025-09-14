package videostreamer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func VideoStreamHandler(streamer VideoStreamer, videoKey string, totalSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rangeHeader := c.GetHeader("Range")
		var start, end int
		var err error

		if rangeHeader == "" {
			start = 0
			end = 1024*1024 - 1
		} else {
			val := strings.TrimPrefix(rangeHeader, "bytes=")
			if strings.HasPrefix(val, "-") {
				suffixLen, convErr := strconv.Atoi(strings.TrimPrefix(val, "-"))
				if convErr != nil || suffixLen <= 0 {
					c.String(http.StatusBadRequest, "Invalid Range")
					return
				}
				if suffixLen > totalSize {
					suffixLen = totalSize
				}
				start = totalSize - suffixLen
				end = totalSize - 1
			} else {
				parts := strings.SplitN(val, "-", 2)
				start, err = strconv.Atoi(parts[0])
				if err != nil || start < 0 {
					c.String(http.StatusBadRequest, "Invalid start byte")
					return
				}
				if len(parts) == 2 && parts[1] != "" {
					end, err = strconv.Atoi(parts[1])
					if err != nil || end < start {
						c.String(http.StatusBadRequest, "Invalid end byte")
						return
					}
				} else {
					end = start + 1024*1024 - 1
				}
			}
		}

		if start >= totalSize {
			c.Header("Content-Range", fmt.Sprintf("bytes */%d", totalSize))
			c.Status(http.StatusRequestedRangeNotSatisfiable)
			return
		}
		if end >= totalSize {
			end = totalSize - 1
		}

		chunk, err := streamer.Seek(videoKey, start, end)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error retrieving video: %v", err)
			return
		}

		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, totalSize))
		c.Header("Accept-Ranges", "bytes")
		c.Header("Content-Length", fmt.Sprintf("%d", len(chunk)))
		c.Header("Content-Type", "video/mp4")
		c.Status(http.StatusPartialContent)
		_, _ = c.Writer.Write(chunk)
	}
}
