package dl

import (
	"envchecker/utils"
	"runtime"
)

func Run(urls []string) {
	ndl := NewDownloader(runtime.NumCPU(), false)
	for _, url := range urls {
		ndl.Download(url, utils.GetFilenameFromUrl(url))
	}
}
