package parser

import (
	"github.com/dominikbraun/espresso/entity"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type (
	metadata map[string]interface{}
	assigner func(val interface{})
)

func readMetadata(metadata metadata, page *entity.Page) {
	mapPrimitive(metadata, "Title", func(val interface{}) {
		page.Title = val.(string)
	})
	mapPrimitive(metadata, "Author", func(val interface{}) {
		page.Author = val.(string)
	})
	mapDate(metadata, "Date", func(val interface{}) {
		page.Date = val.(time.Time)
	})
	mapList(metadata, "Tags", func(val interface{}) {
		page.Tags = append(page.Tags, val.(string))
	})
	mapPrimitive(metadata, "Description", func(val interface{}) {
		page.Description = val.(string)
	})
	mapList(metadata, "Related", func(val interface{}) {
		page.RelatedFQNs = append(page.RelatedFQNs, val.(entity.FQN))
	})
	mapPrimitive(metadata, "Template", func(val interface{}) {
		page.Template = val.(string)
	})
	mapPrimitive(metadata, "Hide", func(val interface{}) {
		page.Hide = val.(bool)
	})
}

func mapPrimitive(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field != nil {
		assigner(field)
	}
}

func mapDate(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	date, err := time.Parse(dateFormat, metadata[key].(string))
	if err != nil {
		panic(err)
	}

	assigner(date)
}

func mapList(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	list, _ := metadata[key].([]interface{})

	for i := 0; i < len(list); i++ {
		assigner(list[i])
	}
}
