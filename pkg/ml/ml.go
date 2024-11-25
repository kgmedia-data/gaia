package ml

type ISummaryML interface {
	BatchSummarize(language string, minSentences, maxSentences int, input []Summary) ([]Summary, error)
}
