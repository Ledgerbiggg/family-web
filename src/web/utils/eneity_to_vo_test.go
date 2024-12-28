package utils

import (
	"fmt"
	"testing"
	"time"
)

type Tag struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        int        `json:"type"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TagVO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

func TestEntityToVO(t *testing.T) {
	now := time.Now()
	tags := []*Tag{
		{Id: 1, Name: "Tag1", Description: "Description1", Type: 1, UpdatedAt: &now},
		{Id: 2, Name: "Tag2", Description: "Description2", Type: 2, UpdatedAt: nil},
	}

	var tagVOs []*TagVO

	err := EntityToVO(&tags, &tagVOs)
	if err != nil {
		t.Errorf("Error: %v", err)
	} else {
		for _, vo := range tagVOs {
			fmt.Printf("Converted VO: %+v\n", vo)
		}
	}
}
