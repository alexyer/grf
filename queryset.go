package grf

type QuerySet interface {
	All() ResponseData
	One(id string) ResponseData
}

type NilQuerySet struct {}

func (q *NilQuerySet) All() *ResponseData {
	return nil
}

func (q *NilQuerySet) One(id string) *ResponseData {
	return nil
}