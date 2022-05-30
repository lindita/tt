package app

import (
	"tt.com/tt/api/app/h5"
	v1 "tt.com/tt/api/app/v1"
)

type App struct {
	H5 *h5.H5
	V1 *v1.V1
}