// Package parser provides a public interface for parsing files
// and converting them to Espresso entities.
package parser

import (
	"github.com/dominikbraun/espresso/entity"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type (
	// metadata represents a map which contains YAML keys and
	// their corresponding values. As those values may be strings
	// or lists, metadata stores them as `interface{}`.
	metadata map[string]interface{}

	// assigner is a function that assigns a given value `val`
	// to a struct field. The field is an enclosed value, so
	// assigner has be be used as a closure.
	//
	// assigners are used by the `readXxx` functions: They read
	// a value from a metadata map and pass that value to the
	// assigner, which is then responsible for assigning that
	// value to an enclosed struct field.
	assigner func(val interface{})
)

// readMetadata takes a metadata map and populates an Espresso
// page with that metadata. The page instance and its fields are
// safe to use after readMetadata has finished.
func readMetadata(metadata metadata, page *entity.Page) {
	readPrimitive(metadata, "Title", func(val interface{}) {
		page.Title = val.(string)
	})

	readPrimitive(metadata, "Author", func(val interface{}) {
		page.Author = val.(string)
	})

	readDate(metadata, "Date", func(val interface{}) {
		page.Date = val.(time.Time)
	})

	readList(metadata, "Tags", func(val interface{}) {
		page.Tags = append(page.Tags, val.(string))
	})

	readPrimitive(metadata, "Description", func(val interface{}) {
		page.Description = val.(string)
	})

	readList(metadata, "Related", func(val interface{}) {
		page.RelatedFQNs = append(page.RelatedFQNs, val.(entity.FQN))
	})

	readPrimitive(metadata, "Template", func(val interface{}) {
		page.Template = val.(string)
	})

	readPrimitive(metadata, "Hide", func(val interface{}) {
		page.Hide = val.(bool)
	})
}

// readPrimitive reads a primitive value with the provided key
// from the metadata map. If the value is valid, readPrimitive
// passes that primitive value to the assigner function.
func readPrimitive(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field != nil {
		assigner(field)
	}
}

// readDate reads a date value with the provided key from the
// metadata map. If the value is valid, readDate passes that
// date value to the assigner function.
func readDate(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	date, err := time.Parse(dateFormat, metadata[key].(string))
	if err != nil {
		panic(err)
	}

	assigner(date)
}

// readList reads a list value with the provided key from the
// metadata map. If the value is valid, readList passes each
// list item to the assigner function.
func readList(metadata metadata, key string, assigner assigner) {
	if field := metadata[key]; field == nil {
		return
	}

	list, _ := metadata[key].([]interface{})

	for i := 0; i < len(list); i++ {
		assigner(list[i])
	}
}
