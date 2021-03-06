// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "Task": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/BcRikko/learning-goa/design
// --out=$(GOPATH)/src/github.com/BcRikko/learning-goa
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// タスク (default view)
//
// Identifier: application/x-learning-goa+json; view=default
type XLearningGoa struct {
	// タスクの作成日時
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// タスクの完了状態
	Done bool `form:"done" json:"done" xml:"done"`
	// タスクID
	ID int `form:"id" json:"id" xml:"id"`
	// タスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
	// タスクの更新日時
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// Validate validates the XLearningGoa media type instance.
func (mt *XLearningGoa) Validate() (err error) {

	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// DecodeXLearningGoa decodes the XLearningGoa instance encoded in resp body.
func (c *Client) DecodeXLearningGoa(resp *http.Response) (*XLearningGoa, error) {
	var decoded XLearningGoa
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// X-Learning-GoaCollection is the media type for an array of X-Learning-Goa (default view)
//
// Identifier: application/x-learning-goa+json; type=collection; view=default
type XLearningGoaCollection []*XLearningGoa

// Validate validates the XLearningGoaCollection media type instance.
func (mt XLearningGoaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeXLearningGoaCollection decodes the XLearningGoaCollection instance encoded in resp body.
func (c *Client) DecodeXLearningGoaCollection(resp *http.Response) (XLearningGoaCollection, error) {
	var decoded XLearningGoaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
