package formatter

import (
	jsoniter "github.com/json-iterator/go"
)

type JSON struct{}

var jsoniterCfg jsoniter.API

func init() {
	jsoniterCfg = jsoniter.Config{SortMapKeys: true}.Froze()
}

func (j *JSON) Format(event *LogEvent) ([]byte, error) {
	data := make(map[string]interface{})
	if label, ok := event.Metadata["label"]; ok {
		if label != "" {
			data["level"] = label
			delete(event.Metadata, "label")
		}
	}
	for k, v := range event.Metadata {
		data[k] = v
	}
	data["msg"] = event.Message
	return jsoniterCfg.Marshal(data)
}
