package main

import (
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestCopyDigits(t *testing.T) {
	var createTempFile = func(t testing.TB, content []byte) (string, error) {
		t.Helper()

		tmpfile, err := os.CreateTemp("", "example")

		if err != nil {
			t.Fatal(err)
		}

		if _, err := tmpfile.Write(content); err != nil {
			t.Fatal(err)
		}

		if err := tmpfile.Close(); err != nil {
			t.Fatal(err)
		}

		return tmpfile.Name(), nil
	}

	var checkGC = func(t testing.TB) {
		t.Helper()

		// Capture memory statistics before and after GC
		var memStatsBefore, memStatsAfter runtime.MemStats
		runtime.ReadMemStats(&memStatsBefore)

		time.Sleep(500 * time.Millisecond)

		// Force garbage collection
		runtime.GC()

		runtime.ReadMemStats(&memStatsAfter)

		t.Logf("Memory stats before GC: Alloc = %v, TotalAlloc = %v, Sys = %v, NumGC = %v",
			memStatsBefore.Alloc, memStatsBefore.TotalAlloc, memStatsBefore.Sys, memStatsBefore.NumGC)
		t.Logf("Memory stats after GC: Alloc = %v, TotalAlloc = %v, Sys = %v, NumGC = %v",
			memStatsAfter.Alloc, memStatsAfter.TotalAlloc, memStatsAfter.Sys, memStatsAfter.NumGC)

		// Check if memory allocation has decreased
		if memStatsAfter.Alloc >= memStatsBefore.Alloc {
			t.Errorf("Expected memory allocation to decrease after GC, but it did not: before = %v, after = %v",
				memStatsBefore.Alloc, memStatsAfter.Alloc)
		}
	}

	t.Run("should get all numbers from file", func(t *testing.T) {
		content := []byte("abc123def456")
		expected := []byte("123456")

		fileName, err := createTempFile(t, content)

		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(fileName) // Clean up after the test

		result := CopyDigits(fileName)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("CopyDigits() = %v, want %v", result, expected)
		}

		checkGC(t)

	})

	t.Run("should work if there are no numbers in file", func(t *testing.T) {
		content := []byte("abcdef")
		expected := []byte("")

		fileName, err := createTempFile(t, content)

		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(fileName) // Clean up after the test

		result := CopyDigits(fileName)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("CopyDigits() = %v, want %v", result, expected)
		}

		checkGC(t)
	})
}
