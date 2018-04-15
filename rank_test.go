package main

import (
	libs "github.com/sobamchan/document-ranking-api/libs"
	"testing"
)

func TestRank(t *testing.T) {
	d1 := Document{1, "i am a dog .", libs.Vector{}}
	d2 := Document{2, "i like small dog .", libs.Vector{}}
	// d3 := Document{3, "i hate cats .", libs.Vector{}}
	d3 := Document{3, "i am a dog .", libs.Vector{}}
	d4 := Document{4, "i have a dog .", libs.Vector{}}
	result := Rank([]Document{d1, d2, d3, d4})
	println("after")
	result.ShowAll()
}
