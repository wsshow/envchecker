package dl

import (
	"envchecker/utils"
	"runtime"
)

func BatchDownload(urls []string) error {
	ndl := NewDownloader(runtime.NumCPU(), false)
	for _, url := range urls {
		err := ndl.Download(url, utils.GetFilenameFromUrl(url))
		if err != nil {
			return err
		}
	}
	return nil
}

func SingleDownload(url string) error {
	ndl := NewDownloader(runtime.NumCPU(), false)
	return ndl.Download(url, utils.GetFilenameFromUrl(url))
}
