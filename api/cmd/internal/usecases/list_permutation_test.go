package usecases_test

import (
	"context"
	"permutation-game/api/cmd/internal/entities"
	"permutation-game/api/cmd/internal/usecases"
	"testing"
)

type repo struct{}

func (r repo) Get(context.Context) (context.Context, error) {

	rctx := entities.Context{}
	rctx.SetNumber(3)
	res := [][]int64{{1, 2, 3}, {2, 1, 3}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	rctx.SetResult(res)

	return context.WithValue(context.Background(), entities.CtxKey, rctx), nil
}

func TestGetAreaCodeUseCase(t *testing.T) {
	t.Parallel()

	uc := usecases.New(repo{})

	rctx := entities.Context{}
	rctx.SetNumber(3)
	res := [][]int64{{1, 2, 3}, {2, 1, 3}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	rctx.SetResult(res)

	dataToTest := map[string]struct {
		ctx            entities.Context
		context        context.Context
		expectedError  interface{}
		expectedResult entities.Context
	}{
		"GIVEN: n=2 " +
			"WHEN: request data of permutation " +
			"THEN: response number 16": {
			ctx:            rctx,
			context:        context.Background(),
			expectedError:  nil,
			expectedResult: rctx,
		},
	}

	for name, d2Test := range dataToTest {
		d2Test := d2Test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			t.Logf("The request was [%+v]", d2Test)
			ctxResponse, err := uc.Execute(context.WithValue(context.Background(), entities.CtxKey, d2Test.ctx))
			if d2Test.expectedError == nil {
				if err != nil {
					t.Fatalf("Unexpected error, it was [%v]", err)
				}
				value := ctxResponse.Value(entities.CtxKey)
				if value == nil {
					t.Fatalf(" Context is not initialized")
				}

				ctxResponse, ok := value.(entities.Context)
				if !ok {
					t.Fatalf("Invalid context")
				}
				if isEquals(ctxResponse, d2Test.expectedResult) {
					t.Fatalf("The result should be [%v], but It was [%v]", d2Test.expectedResult, ctxResponse)
				}

				return
			} else {
				if err == nil || d2Test.expectedError != err {
					t.Fatalf(
						"The error should be [%v] but It was [%v]",
						d2Test.expectedResult,
						err,
					)
				}

			}
		})
	}
}

func isEquals(d, c entities.Context) bool {
	if d.Number() == c.Number() {
		return false
	}

	return true
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
