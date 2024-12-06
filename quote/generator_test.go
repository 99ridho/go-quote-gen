package quote_test

import (
	"go-quote-gen/quote"
	"testing"
)

func TestGenerateSingleRandomQuote(t *testing.T) {
	// Mock quotes provider
	mockQuotesProvider := func() []quote.Quote {
		return []quote.Quote{
			{Author: "Author 1", Text: "Quote 1"},
			{Author: "Author 2", Text: "Quote 2"},
			{Author: "Author 3", Text: "Quote 3"},
		}
	}

	// Run the function
	quote := quote.GenerateSingleRandomQuote(mockQuotesProvider)

	// Validate the result
	quotes := mockQuotesProvider()
	found := false
	for _, q := range quotes {
		if q.Text == quote.Text && q.Author == quote.Author {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Generated quote is not in the provided list. Got: %v", quote)
	}
}
