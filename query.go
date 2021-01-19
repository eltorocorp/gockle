package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// QueryAPI allows gomock mock of gocql.Query
type Query interface {
	Bind(...interface{}) Query
	Exec() error
	Iter() IterAPI
	Scan(...interface{}) error
	ScanCAS(...interface{}) (bool, error)
}

type QueryMock struct {
	mock.Mock
}

func (m QueryMock) Bind(v ...interface{}) QueryMock {
	return m.Called(v).Get(0).(*QueryMock)
}

func (m QueryMock) Exec() error {
	return m.Called().Error(0)
}

func (m QueryMock) Iter() IteratorMock {
	return m.Called().Get(0).(IteratorMock)
}

func (m QueryMock) Scan(values ...interface{}) error {
	return m.Called(values).Error(0)
}

func (m QueryMock) ScanCAS(values ...interface{}) (bool, error) {
	request := m.Called(values)
	return request.Bool(0), request.Error(1)
}

// Query is a wrapper for a query for mockability.
type query struct {
	query *gocql.Query
}

// NewQuery instantiates a new Query
func NewQuery(query *gocql.Query) QueryAPI {
	return &query{
		query,
	}
}

// Bind wraps the query's Bind method
func (q *query) Bind(v ...interface{}) QueryAPI {
	return NewQuery(q.query.Bind(v...))
}

// Exec wraps the query's Exec method
func (q *query) Exec() error {
	return q.query.Exec()
}

// Iter wraps the query's Iter method
func (q *query) Iter() IterAPI {
	return NewIter(q.query.Iter())
}

// Scan wraps the query's Scan method
func (q *query) Scan(dest ...interface{}) error {
	return q.query.Scan(dest...)
}

func (q *query) ScanCAS(dest ...interface{}) (bool, error) {
	return q.query.ScanCAS(dest)
}
