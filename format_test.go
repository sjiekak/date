// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date_test

import (
	"testing"
	"time"

	"github.com/sjiekak/date"
)

func TestParseISO(t *testing.T) {
	cases := []struct {
		value string
		year  int
		month time.Month
		day   int
	}{
		{"1969-12-31", 1969, time.December, 31},
		{"+1970-01-01", 1970, time.January, 1},
		{"+01970-01-02", 1970, time.January, 2},
		{"2000-02-28", 2000, time.February, 28},
		{"+2000-02-29", 2000, time.February, 29},
		{"+02000-03-01", 2000, time.March, 1},
		{"+002004-02-28", 2004, time.February, 28},
		{"2004-02-29", 2004, time.February, 29},
		{"2004-03-01", 2004, time.March, 1},
		{"0000-01-01", 0, time.January, 1},
		{"+0001-02-03", 1, time.February, 3},
		{"+00019-03-04", 19, time.March, 4},
		{"0100-04-05", 100, time.April, 5},
		{"2000-05-06", 2000, time.May, 6},
		{"+30000-06-07", 30000, time.June, 7},
		{"+400000-07-08", 400000, time.July, 8},
		{"+5000000-08-09", 5000000, time.August, 9},
		{"-0001-09-11", -1, time.September, 11},
		{"-0019-10-12", -19, time.October, 12},
		{"-00100-11-13", -100, time.November, 13},
		{"-02000-12-14", -2000, time.December, 14},
		{"-30000-02-15", -30000, time.February, 15},
		{"-0400000-05-16", -400000, time.May, 16},
		{"-5000000-09-17", -5000000, time.September, 17},
	}
	for _, c := range cases {
		d, err := date.ParseISO(c.value)
		if err != nil {
			t.Errorf("ParseISO(%v) == %v", c.value, err)
		}
		year, month, day := d.Date()
		if year != c.year || month != c.month || day != c.day {
			t.Errorf("ParseISO(%v) == %v, want (%v, %v, %v)", c.value, d, c.year, c.month, c.day)
		}
	}

	badCases := []string{
		"1234-05",
		"1234-5-6",
		"1234-05-6",
		"1234-5-06",
		"12340506",
		"1234/05/06",
		"1234-0A-06",
		"1234-05-0B",
		"1234-05-06trailing",
		"padding1234-05-06",
		"1-02-03",
		"10-11-12",
		"100-02-03",
		"+1-02-03",
		"+10-11-12",
		"+100-02-03",
		"-123-05-06",
	}
	for _, c := range badCases {
		d, err := date.ParseISO(c)
		if err == nil {
			t.Errorf("ParseISO(%v) == %v", c, d)
		}
	}
}

func TestParse(t *testing.T) {
	// Test ability to parse a few common date formats
	cases := []struct {
		layout string
		value  string
		year   int
		month  time.Month
		day    int
	}{
		{date.ISO8601, "1969-12-31", 1969, time.December, 31},
		{date.ISO8601B, "19700101", 1970, time.January, 1},
		{date.RFC822, "29-Feb-00", 2000, time.February, 29},
		{date.RFC822W, "Mon, 01-Mar-04", 2004, time.March, 1},
		{date.RFC850, "Wednesday, 12-Aug-15", 2015, time.August, 12},
		{date.RFC1123, "05 Dec 1928", 1928, time.December, 5},
		{date.RFC1123W, "Mon, 05 Dec 1928", 1928, time.December, 5},
		{date.RFC3339, "2345-06-07", 2345, time.June, 7},
	}
	for _, c := range cases {
		d, err := date.Parse(c.layout, c.value)
		if err != nil {
			t.Errorf("Parse(%v) == %v", c.value, err)
		}
		year, month, day := d.Date()
		if year != c.year || month != c.month || day != c.day {
			t.Errorf("Parse(%v) == %v, want (%v, %v, %v)", c.value, d, c.year, c.month, c.day)
		}
	}

	// Test inability to parse ISO 8601 expanded year format
	badCases := []string{
		"+1234-05-06",
		"+12345-06-07",
		"-1234-05-06",
		"-12345-06-07",
	}
	for _, c := range badCases {
		d, err := date.Parse(date.ISO8601, c)
		if err == nil {
			t.Errorf("Parse(%v) == %v", c, d)
		}
	}
}

func TestFormatISO(t *testing.T) {
	cases := []struct {
		value string
		n     int
	}{
		{"-5000-02-03", 4},
		{"-05000-02-03", 5},
		{"-005000-02-03", 6},
		{"+0000-01-01", 4},
		{"+00000-01-01", 5},
		{"+1000-01-01", 4},
		{"+01000-01-01", 5},
		{"+1970-01-01", 4},
		{"+001999-12-31", 6},
		{"+999999-12-31", 6},
	}
	for _, c := range cases {
		d, err := date.ParseISO(c.value)
		if err != nil {
			t.Errorf("FormatISO(%v) cannot parse input: %v", c.value, err)
			continue
		}
		value := d.FormatISO(c.n)
		if value != c.value {
			t.Errorf("FormatISO(%v) == %v, want %v", c, value, c.value)
		}
	}
}
