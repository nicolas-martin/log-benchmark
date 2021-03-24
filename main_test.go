package main

import (
	"testing"

	"github.com/nicolas-martin/coord-gen/log-benchmark/hooks"
	"github.com/nicolas-martin/coord-gen/log-benchmark/regex"
	log "github.com/sirupsen/logrus"
)

func benchmarkHooks(level log.Level, b *testing.B) {
	l := hooks.Setup()
	for n := 0; n < b.N; n++ {
		l.Log(level, "hello hooks")
	}
}

func BenchmarkHooks_Info(b *testing.B) {
	benchmarkHooks(log.InfoLevel, b)
}

func BenchmarkHooks_Err(b *testing.B) {
	benchmarkHooks(log.ErrorLevel, b)
}

func benchmarkRegex(level log.Level, b *testing.B) {
	l := regex.Setup()
	for n := 0; n < b.N; n++ {
		l.Log(level, "hello regex")
	}
}

func BenchmarkRegex_Info(b *testing.B) {
	benchmarkRegex(log.InfoLevel, b)
}

func BenchmarkRegex_Err(b *testing.B) {
	benchmarkRegex(log.ErrorLevel, b)
}
