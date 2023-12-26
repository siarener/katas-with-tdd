package kata

import (
	"reflect"
	"testing"
)

func TestBalancedParentheses(t *testing.T) {
	t.Run("no pair", func(t *testing.T) {
		assertSameParentheses(t, BalancedParentheses(0), []string{""})
	})
	t.Run("one pair", func(t *testing.T) {
		assertSameParentheses(t, BalancedParentheses(1), []string{"()"})
	})
	t.Run("two pairs", func(t *testing.T) {
		assertSameParentheses(t, BalancedParentheses(2), []string{"(())", "()()"})
	})
	t.Run("three pairs", func(t *testing.T) {
		assertSameParentheses(t, BalancedParentheses(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	})
	t.Run("four pairs", func(t *testing.T) {
		assertSameParentheses(t, BalancedParentheses(4), []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"})
	})
}

func bothSlicesAreEmpty(a, b []string) bool {
	return len(a) == 0 && len(b) == 0
}

func assertSameParentheses(t *testing.T, got, want []string) {
	if !reflect.DeepEqual(got, want) && !bothSlicesAreEmpty(got, want) {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}
