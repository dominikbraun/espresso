package parser

import (
	"github.com/dominikbraun/espresso/entity"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type (
	metadata  map[string]interface{}
	assigner  func(val interface{})
	processor func(metadata metadata, key string, assigner assigner)
	mapping   struct {
		key       string
		processor processor
		assigner  assigner
	}
)

func readMetadata(metadata metadata, page *entity.Page) {
	mappings := []mapping{
		{
			key:       "Title",
			processor: processPrimitive,
			assigner: func(val interface{}) {
				page.Title = val.(string)
			},
		},
		{
			key:       "Author",
			processor: processPrimitive,
			assigner: func(val interface{}) {
				page.Author = val.(string)
			},
		},
		{
			key:       "Date",
			processor: processDate,
			assigner: func(val interface{}) {
				page.Date = val.(time.Time)
			},
		},
		{
			key:       "Tags",
			processor: processList,
			assigner: func(val interface{}) {
				page.Tags = append(page.Tags, val.(string))
			},
		},
		{
			key:       "Description",
			processor: processPrimitive,
			assigner: func(val interface{}) {
				page.Description = val.(string)
			},
		},
		{
			key:       "Related",
			processor: processList,
			assigner: func(val interface{}) {
				page.RelatedFQNs = append(page.RelatedFQNs, val.(entity.FQN))
			},
		},
		{
			key:       "Template",
			processor: processPrimitive,
			assigner: func(val interface{}) {
				page.Template = val.(string)
			},
		},
		{
			key:       "Hide",
			processor: processPrimitive,
			assigner: func(val interface{}) {
				page.Hide = val.(bool)
			},
		},
	}

	for _, m := range mappings {
		m.processor(metadata, m.key, m.assigner)
	}
}

func processPrimitive(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field != nil {
		assigner(field)
	}
}

func processDate(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	date, err := time.Parse(dateFormat, metadata[key].(string))
	if err != nil {
		panic(err)
	}

	assigner(date)
}

func processList(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	list, _ := metadata[key].([]interface{})

	for i := 0; i < len(list); i++ {
		assigner(list[i])
	}
}
