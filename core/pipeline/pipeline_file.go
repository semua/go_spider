package pipeline

import (
	"fmt"
	"os"

	"github.com/semua/go_spider/core/common/com_interfaces"
	"github.com/semua/go_spider/core/common/page_items"
)

type PipelineFile struct {
	pFile *os.File

	path string
}

func NewPipelineFile(path string) *PipelineFile {
	pFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("File '" + path + "' in PipelineFile open failed.")
	}
	return &PipelineFile{path: path, pFile: pFile}
}

func (this *PipelineFile) Process(items *page_items.PageItems, t com_interfaces.Task) {
	this.pFile.WriteString("----------------------------------------------------------------------------------------------\n")
	this.pFile.WriteString("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	this.pFile.WriteString("Crawled result : \n")
	for key, value := range items.GetAll() {
		switch value.(type) {
		case []string:
			this.pFile.WriteString(key + "\t:\t" + fmt.Sprint(value) + "\n")
			break
		case map[string]string:
			this.pFile.WriteString(key + "\t:\t" + fmt.Sprint(value) + "\n")
			break
		case string:
			this.pFile.WriteString(key + "\t:\t" + value.(string) + "\n")
			break
		default:
			panic("type match miss")
		}
	}
}
