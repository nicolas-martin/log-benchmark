package main

import (
	"fmt"
	"testing"

	"github.com/nicolas-martin/log-benchmark/hooks"
	"github.com/nicolas-martin/log-benchmark/regex"
	"github.com/sirupsen/logrus"
)

func main() {
	resHookInfo := testing.Benchmark(BenchmarkHooksInfo)
	resHookErr := testing.Benchmark(BenchmarkHooksErr)
	resRegexInfo := testing.Benchmark(BenchmarkRegexInfo)
	resRegexErr := testing.Benchmark(BenchmarkRegexErr)
	resNormalInfo := testing.Benchmark(BenchmarkNormalInfo)
	resNormalErr := testing.Benchmark(BenchmarkNormalErr)

	fmt.Printf("HookInfo\t %s\n", resHookInfo)
	fmt.Printf("HookErr\t\t %s\n", resHookErr)
	fmt.Printf("RegexInfo\t %s\n", resRegexInfo)
	fmt.Printf("RegexErr\t %s\n", resRegexErr)
	fmt.Printf("NormalInfo\t %s\n", resNormalInfo)
	fmt.Printf("NormalErr\t %s\n", resNormalErr)
}

func benchmarkLog(l *logrus.Logger, level logrus.Level, b *testing.B, msg string) {
	for n := 0; n < b.N; n++ {
		l.Log(level, msg)
	}
}

func BenchmarkHooksInfo(b *testing.B) {
	l := hooks.Setup()
	benchmarkLog(l, logrus.InfoLevel, b, "Info Hook")
}

func BenchmarkHooksErr(b *testing.B) {
	l := hooks.Setup()
	benchmarkLog(l, logrus.ErrorLevel, b, "Err Hook")
}

func BenchmarkRegexInfo(b *testing.B) {
	l := regex.Setup()
	benchmarkLog(l, logrus.InfoLevel, b, "Info Regex")
}

func BenchmarkRegexErr(b *testing.B) {
	l := regex.Setup()
	benchmarkLog(l, logrus.ErrorLevel, b, "Err Regex")
}

func BenchmarkNormalInfo(b *testing.B) {
	l := logrus.New()
	benchmarkLog(l, logrus.InfoLevel, b, "Info Normal")
}

func BenchmarkNormalErr(b *testing.B) {
	l := logrus.New()
	benchmarkLog(l, logrus.ErrorLevel, b, "Err Normal")
}
