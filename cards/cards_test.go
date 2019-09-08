package cards

import (
    "testing"
)

func TestNewDeck(t *testing.T) {
    actual := NewDeck().Cards[0].String()
    expected := "1s"
    if actual != expected {
        t.Errorf("got %v. expected %v",
        actual, expected)
    }

    actual = NewDeck().Cards[12].String()
    expected = "Ks"
    if actual != expected {
        t.Errorf("got %v. expected %v",
        actual, expected)
    }

    actual = NewDeck().Cards[13].String()
    expected = "1h"
    if actual != expected {
        t.Errorf("got %v. expected %v",
        actual, expected)
    }
}
