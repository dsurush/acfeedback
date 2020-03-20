package models

type FeedBack struct {
	Id           int64  `json:"id"`
	Feedback     string `json:"feedback,omitempty"`
	Who          int64 `json:"who,omitempty"`
	Whom         int64 `json:"whom,omitempty"`
	FeedbackTime string `json:"feedbackTime"`
	Remove       bool   `json:"remove"`
}
