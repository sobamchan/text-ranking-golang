package main

import (
	libs "github.com/sobamchan/document-ranking-api/libs"
	"sort"
	// libs "./libs"
)

type Document struct {
	ID          int
	Content     string
	TFIDFVector libs.Vector
}

type Candidate struct {
	ID       int
	TFIDF    libs.TfidfResult
	Distance libs.Distance
}

type Candidates []Candidate

func (c Candidates) Len() int {
	return len(c)
}

func (c Candidates) Less(i, j int) bool {
	return c[i].Distance < c[j].Distance
}

func (c Candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Candidates) ShowAll() {
	for _, candidate := range c {
		println(candidate.ID)
	}
}

type Documents []Document

func Rank(documents Documents) Candidates {

	// id2idx := map[int]int{}
	idx2id := map[int]int{}
	for i, d := range documents[1:] {
		idx2id[i] = d.ID
		// id2idx[d.ID] = i
	}

	var contents []string
	for _, document := range documents {
		contents = append(contents, document.Content)
	}
	Tfidf := libs.TFIDFNew(contents)
	TfidfResults := Tfidf.Fit()

	target := TfidfResults[0]
	var candidates Candidates
	for i, tfdifResult := range TfidfResults[1:] {
		candidates = append(candidates,
			Candidate{idx2id[i], tfdifResult, 0.0})
	}

	for i, candidate := range candidates {
		distance := libs.CosineDistance(libs.Vector(target),
			libs.Vector(candidate.TFIDF))
		candidates[i].Distance = distance
	}

	sort.Sort(candidates)

	return candidates
}
