package main

import (
	"strings"
	"testing"
)

// TestImportPrompt pins the prompt's load-bearing tokens, the way
// coach_test.go and skills_test.go pin theirs.
func TestImportPrompt(t *testing.T) {
	for _, want := range []string{
		"[file:",               // how omni hands over a photo or document the owner sent
		"Owner's message",      // how omni hands over the user's trailing words
		"pecunia_accounts",     // the targets a file maps onto
		"pecunia_credit_cards", //
		"pecunia_categories",   // where rows get categorized from
		"pecunia_transactions", // the dedupe read and the write
		"pay_bill",             // a statement payment is not a transaction
		"installments",         // N parcels are one row
		"minor units",          // the amount rule
		"across currencies",    // never sum or compare across currencies
		"photo",                // receipts and screenshots are in scope
		"duplicate",            // the skip-what-is-filed step
		"no tables",            // telegram style
	} {
		if !strings.Contains(importPrompt, want) {
			t.Errorf("import prompt does not contain %q", want)
		}
	}
}
