package api

import "time"

type Emoji struct{}

type DescData struct {
	Emoji Emoji `json:"emoji"`
}

type TrelloAttachmentType struct {
	Board int `json:"board"`
	Card  int `json:"card"`
}

type AttachmentsByType struct {
	Trello TrelloAttachmentType `json:"trello"`
}

type Badges struct {
	AttachmentsByType     AttachmentsByType `json:"attachmentsByType"`
	Location              bool              `json:"location"`
	Votes                 int               `json:"votes"`
	ViewingMemberVoted    bool              `json:"viewingMemberVoted"`
	Subscribed            bool              `json:"subscribed"`
	Fogbugz               string            `json:"fogbugz"`
	CheckItems            int               `json:"checkItems"`
	CheckItemsChecked     int               `json:"checkItemsChecked"`
	CheckItemsEarliestDue interface{}       `json:"checkItemsEarliestDue"`
	Comments              int               `json:"comments"`
	Attachments           int               `json:"attachments"`
	Description           bool              `json:"description"`
	Due                   time.Time         `json:"due"`
	DueComplete           bool              `json:"dueComplete"`
}

type Labels struct {
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type Cover struct {
	IDAttachment         interface{} `json:"idAttachment"`
	Color                interface{} `json:"color"`
	IDUploadedBackground interface{} `json:"idUploadedBackground"`
	Size                 string      `json:"size"`
	Brightness           string      `json:"brightness"`
}

type ListApiResponse []struct {
	ID                    string        `json:"id"`
	CheckItemStates       interface{}   `json:"checkItemStates"`
	Closed                bool          `json:"closed"`
	DateLastActivity      time.Time     `json:"dateLastActivity"`
	Desc                  string        `json:"desc"`
	DescData              DescData      `json:"descData"`
	DueReminder           int           `json:"dueReminder"`
	IDBoard               string        `json:"idBoard"`
	IDList                string        `json:"idList"`
	IDMembersVoted        []interface{} `json:"idMembersVoted"`
	IDShort               int           `json:"idShort"`
	IDAttachmentCover     interface{}   `json:"idAttachmentCover"`
	IDLabels              []string      `json:"idLabels"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Name                  string        `json:"name"`
	Pos                   int           `json:"pos"`
	ShortLink             string        `json:"shortLink"`
	IsTemplate            bool          `json:"isTemplate"`
	Badges                Badges        `json:"badges"`
	DueComplete           bool          `json:"dueComplete"`
	Due                   time.Time     `json:"due"`
	IDChecklists          []interface{} `json:"idChecklists"`
	IDMembers             []interface{} `json:"idMembers"`
	Labels                []Labels      `json:"labels"`
	ShortURL              string        `json:"shortUrl"`
	Subscribed            bool          `json:"subscribed"`
	URL                   string        `json:"url"`
	Cover                 Cover         `json:"cover"`
}
