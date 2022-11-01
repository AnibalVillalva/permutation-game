package entrypoints

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"permutation-game/api/cmd/internal/entities"
	"permutation-game/api/cmd/internal/usecases"
	"testing"

	"github.com/gin-gonic/gin"
)

/*
import (
	"permutation-game/internal/core"
	"permutation-game/internal/core/entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/*
 GIVEN: n=4
 WHEN: request data of permutation
 THEN: response number 16
 and series:
 	[1,2,3,4],
	[2,1,3,4],
	[3,1,2,4],
	[3,2,1,4],
 	[1,2,4,3],
	[2,1,4,3],
	[3,1,4,2],
	[3,2,4,1]
 	[1,4,2,3],
	[2,4,1,3],
	[3,4,1,2],
	[3,4,2,1]
 	[4,1,2,3],
	[4,2,1,3],
	[4,3,1,2],
	[4,3,2,1]
*/
type repo struct{}

func (r repo) Get(c context.Context) (context.Context, error) {
	rctx := entities.Context{}
	value := c.Value(entities.CtxKey)
	m, _ := value.(entities.Context)

	switch m.Number() {
	case 0:
		return nil, errors.New("Cannot build permutation")
	case 2:
		{
			res := [][]int64{{1, 2, 3}, {2, 1, 3}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
			rctx = rctx.SetNumber(3).SetResult(res)
		}
	}

	return context.WithValue(context.Background(), entities.CtxKey, rctx), nil
}

func TestGetAreaCodeUseCase(t *testing.T) {
	t.Parallel()
	uc := usecases.New(repo{})
	e := entrypointsImpl{
		uc: *uc,
	}

	rctx := entities.Context{}
	rctx.SetNumber(3)
	res := [][]int64{{1, 2, 3}, {2, 1, 3}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	rctx.SetResult(res)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c1, _ := gin.CreateTestContext(w)
	c1.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	params1 := url.Values{}
	params1.Set("number", "string")
	c1.Request.URL.RawQuery = params1.Encode()

	c2, _ := gin.CreateTestContext(w)
	c2.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	c3, _ := gin.CreateTestContext(w)
	c3.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	params3 := url.Values{}
	params3.Set("number", "2")
	c3.Request.URL.RawQuery = params3.Encode()

	c4, _ := gin.CreateTestContext(w)
	c4.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	params4 := url.Values{}
	params4.Set("number", "0")
	c4.Request.URL.RawQuery = params4.Encode()

	//c.DefaultQuery("number")
	dataToTest := map[string]struct {
		context        *gin.Context
		expectedError  interface{}
		expectedResult *entities.Context
	}{
		"GIVEN: request " +
			"WHEN: request param number is string " +
			"THEN: response is bad request": {
			context:        c1,
			expectedError:  http.StatusBadRequest,
			expectedResult: &rctx,
		},
		"GIVEN: request " +
			"WHEN: does not initialize params " +
			"THEN: response bad request": {
			context:        c2,
			expectedError:  http.StatusBadRequest,
			expectedResult: &rctx,
		},
		"GIVEN: n=2 " +
			"WHEN: request data of permutation " +
			"THEN: response number 16": {
			context:        c3,
			expectedError:  http.StatusOK,
			expectedResult: &rctx,
		},
		"GIVEN: n=0 " +
			"WHEN: request data of permutation " +
			"THEN: response error": {
			context:        c4,
			expectedError:  http.StatusNotFound,
			expectedResult: &rctx,
		},
	}

	for name, d2Test := range dataToTest {
		d2Test := d2Test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			t.Logf("The request was [%+v]", d2Test)
			e.Execute(d2Test.context)
			if d2Test.expectedError == 0 {
				if w.Code != d2Test.expectedError {
					t.Fatalf("Unexpected error, it was [%d]", w.Code)
				}
				b, _ := ioutil.ReadAll(w.Body)
				// if b == d2Test.expectedResult) {
				// 	t.Fatalf("The result should be [%v], but It was [%v]", d2Test.expectedResult, ctxResponse)
				// }
				t.Logf("response: %+v", b)
				return

			}
		})
	}
}

func TestPing(t *testing.T) {
	uc := usecases.New(repo{})
	e := entrypointsImpl{
		uc: *uc,
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	e.Pong(c)
}

func BenchmarkGetAreaCodeUseCase(b *testing.B) {
	rctx := entities.Context{}
	rctx.SetNumber(3)
	res := [][]int64{{1, 2, 3}, {2, 1, 3}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	rctx.SetResult(res)
	uc := usecases.New(repo{})

	for i := 0; i < b.N; i++ {
		_, _ = uc.Execute(context.WithValue(context.Background(), entities.CtxKey, rctx))
	}
}
