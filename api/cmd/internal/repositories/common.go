package repositories

//go:generate mockgen -source=${GOFILE} -destination=$PWD/internal/mocks/repositories/$GOPACKAGE/mock_${GOFILE} -package=mocks

import (
	"context"
)

// Getter repo Interface.
type Getter interface {
	Get(context.Context) (context.Context, error)
}
