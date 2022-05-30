package app

import (
	"tt.com/tt/api/app/h5"
	v1 "tt.com/tt/api/app/v1"
)

type App struct {
	H5 *AppH5
	V1 *AppV1
}

type AppV1 struct {
	Tt *v1.TtController
}

type AppH5 struct {
	Tt *h5.TtController
}