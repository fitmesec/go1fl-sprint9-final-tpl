package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("positive size", func(t *testing.T) {
		size := 10
		data := generateRandomElements(size)

		if len(data) != size {
			t.Errorf("ожидалась длина слайса %d, получено %d", size, len(data))
		}

		for _, v := range data {
			if v <= 0 {
				t.Errorf("ожидалось положительное число, получено %d", v)
			}
		}
	})

	t.Run("zero size", func(t *testing.T) {
		data := generateRandomElements(0)

		if len(data) != 0 {
			t.Errorf("ожидался пустой слайс, получено %d элементов", len(data))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		data := generateRandomElements(-5)

		if len(data) != 0 {
			t.Errorf("ожидался пустой слайс, получено %d элементов", len(data))
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("regular slice", func(t *testing.T) {
		data := []int{3, 7, 2, 9, 4}
		want := 9

		got := maximum(data)
		if got != want {
			t.Errorf("ожидалось %d, получено %d", want, got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		data := []int{42}
		want := 42

		got := maximum(data)
		if got != want {
			t.Errorf("ожидалось %d, получено %d", want, got)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		data := []int{}
		want := 0

		got := maximum(data)
		if got != want {
			t.Errorf("ожидалось %d, получено %d", want, got)
		}
	})

	t.Run("all equal elements", func(t *testing.T) {
		data := []int{5, 5, 5, 5}
		want := 5

		got := maximum(data)
		if got != want {
			t.Errorf("ожидалось %d, получено %d", want, got)
		}
	})
}
