package ba_baracus

import (
	mr_t "github.com/tjarratt/mr_t"
)

type Baracus interface {
	N int
	TestingT

	ReportAllocs()
	ResetTimer()
	SetBytes(n int64)
	StartTimer()
	StopTimer()
}

type BaBaracus struct {
	TestingT
}

func (b BaBaracus) ReportAllocs() { }

func (b BaBaracus) ResetTimer() { }

func (b BaBaracus) SetBytes(n int64) { }

func (b BaBaracus) StartTimer() {}

func (b BaBaracus) StopTimer() { }
