package main

type DseOutSpcObject struct {
	Langs DseOutLanguages   `json:"languages"`
	Spc   []PathCondSegment `json:"segmented_path_condition"`
}

type DseOutLanguages struct {
	Smt string `json:"smt"`
}

type PathCondSegment struct {
	Event   *string        `json:"callback"`
	Segment []PathCondItem `json:"path_cond_segment"`
}

type PathCondItem struct {
	Constraint string `json:"constraint"`
	Followed   bool   `json:"followed_value"`
}
