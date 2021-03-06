package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wikisophia/api-arguments/arguments/argumentstest"
	"github.com/wikisophia/api-arguments/arguments/memory"
)

// TestInMemoryStore makes sure that the inMemoryStore is consistent with the StoreTests suite.
// This helps verify:
//    1. The inMemoryStore, which is used throughout app tests to avoid a DB dependency.
//    2.The StoreTests suite, which is reused to test the real Postgres implementation.
func TestInMemoryStore(t *testing.T) {
	suite.Run(t, &InMemoryTests{
		Store: memory.NewStore(),
	})
}

type InMemoryTests argumentstest.StoreTests

func (suite *InMemoryTests) SetupTest() {
	suite.Store = memory.NewStore()
}
