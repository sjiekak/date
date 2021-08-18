package date

import "testing"

func BenchmarkFormatISORegexp(b *testing.B) {
	days := make([]string, 0, b.N)
	day := Today()
	for n := 0; n < b.N; n++ {
		days = append(days, day.String())
		day.Add(1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		parseISORegexp(days[n])
	}
}

func BenchmarkFormatISORune(b *testing.B) {
	days := make([]string, 0, b.N)
	day := Today()
	for n := 0; n < b.N; n++ {
		days = append(days, day.String())
		day.Add(1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		parseISORune(days[n])
	}
}
