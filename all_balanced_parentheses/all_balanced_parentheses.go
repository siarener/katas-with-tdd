/* https://www.codewars.com/kata/5426d7a2c2c7784365000783/train/go

Write a function which makes a list of strings representing all of the ways you can balance n pairs of parentheses
Examples

balancedParens 0 -> [""]
balancedParens 1 -> ["()"]
balancedParens 2 -> ["()()","(())"]
balancedParens 3 -> ["()()()","(())()","()(())","(()())","((()))"] */

package kata

const OPEN_BRACKET = "("
const CLOSE_BRACKET = ")"

func addParentheses(results []string, wipOption string, remainingStarts, remainingEnds int) []string {
	if remainingEnds == 0 {
		return append(results, wipOption)
	} else if remainingEnds == remainingStarts {
		return addParentheses(results, wipOption+OPEN_BRACKET, remainingStarts-1, remainingEnds)
	} else if remainingEnds > remainingStarts {
		if remainingStarts == 0 {
			return addParentheses(results, wipOption+CLOSE_BRACKET, remainingStarts, remainingEnds-1)
		} else {
			results = append(results, addParentheses([]string{}, wipOption+OPEN_BRACKET, remainingStarts-1, remainingEnds)...)
			results = append(results, addParentheses([]string{}, wipOption+CLOSE_BRACKET, remainingStarts, remainingEnds-1)...)
			return results
		}
	}
	return results

}

func BalancedParentheses(n int) []string {
	return addParentheses([]string{}, "", n, n)
}
