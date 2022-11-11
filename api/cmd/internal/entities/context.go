package entities

type key int

const (
	CtxKey key = iota
)

type Context struct {
	aPIError string
	number   int64
	tuples   [][]int64
	order    int64
}

func Build() Context {
	return Context{}
}

func (c Context) SetAPIError(a string) Context {
	c.aPIError = a
	return c
}

func (c Context) SetNumber(n int64) Context {
	c.number = n
	return c
}

func (c Context) SetResult(r [][]int64) Context {
	result := make([][]int64, len(r))
	for i := range r {
		result[i] = make([]int64, len(r[i]))
		copy(result[i], r[i])
	}

	c.result = result

	return c
}

func (c Context) APIError() string {
	return c.aPIError
}

func (c Context) Result() [][]int64 {
	return c.result
}

func (c Context) Number() int64 {
	return c.number
}
