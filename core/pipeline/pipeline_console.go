package pipeline

import (
	"github.com/semua/go_spider/core/common/com_interfaces"
	"github.com/semua/go_spider/core/common/page_items"
)

type PipelineConsole struct {
}

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{}
}

func (this *PipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	println("Crawled result : ")
	for key, value := range items.GetAll() {
		//println(key + "\t:\t" + value)
		switch value.(type) {
		case []string:
			println(key, "\t:\t", value.([]string))
			break
		case map[string]string:
			println(key, "\t:\t", value.(map[string]string))
			break
		case string:
			println(key, "\t:\t", value.(string))
			break
		default:
			panic("type match miss")
		}
	}
}
