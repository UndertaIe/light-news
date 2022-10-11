package main

var dataCheck Checker = new(DataQualityCheck)

type Checker interface {
	Check([]*NewsModel) error
}

type DataQualityCheck struct { //TODO: 接入告警

}

func (dqc *DataQualityCheck) Check(ms []*NewsModel) error {
	for _, m := range ms {
		if m.Title == "" {
			return ErrModelFieldTitleEmpty
		}
		if m.NewsUrl == "" {
			return ErrModelFieldNewsURLEmpty
		}
	}
	return nil
}
