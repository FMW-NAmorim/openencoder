package net

import "github.com/alfg/enc/types"

// DownloadFunc creates a download.
type DownloadFunc func(job types.Job) error

// GetDownloadFunc sets the download function.
func GetDownloadFunc() DownloadFunc {
	return S3Download
}