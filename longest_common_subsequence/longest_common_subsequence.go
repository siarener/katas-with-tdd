/*
https://www.codewars.com/kata/52756e5ad454534f220001ef/train/go

Write a function called LCS that accepts two sequences and returns the longest subsequence common to the passed in sequences.
Subsequence

A subsequence is different from a substring. The terms of a subsequence need not be consecutive terms of the original sequence.
Example subsequence

Subsequences of "abc" = "a", "b", "c", "ab", "ac", "bc" and "abc".
LCS examples

LCS( "abcdef", "abc" ) => returns "abc"
LCS( "abcdef", "acf" ) => returns "acf"
LCS( "132535365", "123456789" ) => returns "12356"

Notes

	    Both arguments will be strings
	    Return value must be a string
	    Return an empty string if there exists no common subsequence
	    Both arguments will have one or more characters (in JavaScript)
	    All tests will only have a single longest common subsequence. Don't worry about cases such as LCS( "1234", "3412" ), which would have two possible longest common subsequences: "12" and "34".

		[...]

# Tips

Wikipedia has an explanation of the two properties that can be used to solve the problem:

	    First property
			[
			LCS(X^A,Y^A) = LCS(X,Y)^A, for all strings X, Y and all symbols A, where ^ denotes string concatenation.
			This allows one to simplify the LCS computation for two sequences ending in the same symbol. For example,
			LCS("BANANA","ATANA") = LCS("BANAN","ATAN")^"A", Continuing for the remaining common symbols,
			LCS("BANANA","ATANA") = LCS("BAN","AT")^"ANA".
			]

	    Second property
			[
			If A and B are distinct symbols (A≠B), then LCS(X^A,Y^B) is one of the maximal-length strings in the set
			{ LCS(X^A,Y), LCS(X,Y^B) }, for all strings X, Y.

			For example, LCS("ABCDEFG","BCDGK") is the longest string among LCS("ABCDEFG","BCDG") and LCS("ABCDEF","BCDGK");
			if both happened to be of equal length, one of them could be chosen arbitrarily.

			To realize the property, distinguish two cases:
	    		* If LCS("ABCDEFG","BCDGK") ends with a "G", then the final "K" cannot be in the LCS,
				  hence LCS("ABCDEFG","BCDGK") = LCS("ABCDEFG","BCDG").
	    		* If LCS("ABCDEFG","BCDGK") does not end with a "G", then the final "G" cannot be in the LCS,
				  hence LCS("ABCDEFG","BCDGK") = LCS("ABCDEF","BCDGK").
			]
*/
package kata

type TmpLcs struct {
	s1               string
	s2               string
	foundSubsequence string
}

func computeLcs(tmpLcs TmpLcs) TmpLcs {

	for tmpLcs.s1 != "" && tmpLcs.s2 != "" {
		lenS1 := len(tmpLcs.s1)
		lenS2 := len(tmpLcs.s2)

		if tmpLcs.s1[lenS1-1] == tmpLcs.s2[lenS2-1] {

			/* use first property:
			LCS(X^A,Y^A) = LCS(X,Y)^A, for all strings X, Y and all symbols A, where ^ denotes string concatenation. */
			tmpLcs = TmpLcs{
				tmpLcs.s1[:lenS1-1],
				tmpLcs.s2[:lenS2-1],
				string(tmpLcs.s1[lenS1-1]) + tmpLcs.foundSubsequence,
			}

		} else {
			/* use second property:
			If A and B are distinct symbols (A≠B), then LCS(X^A,Y^B) is one of the maximal-length strings in the set
			{ LCS(X^A,Y), LCS(X,Y^B) }, for all strings X, Y.*/
			optionA := computeLcs(TmpLcs{tmpLcs.s1[:lenS1-1], tmpLcs.s2, ""})
			optionB := computeLcs(TmpLcs{tmpLcs.s1, tmpLcs.s2[:lenS2-1], ""})

			if len(optionA.foundSubsequence) > len(optionB.foundSubsequence) {
				optionA.foundSubsequence += tmpLcs.foundSubsequence
				tmpLcs = optionA
			} else {
				optionB.foundSubsequence += tmpLcs.foundSubsequence
				tmpLcs = optionB
			}
		}
	}
	return tmpLcs
}

func LCS(s1, s2 string) string {
	tmpLcs := computeLcs(TmpLcs{s1, s2, ""})
	return tmpLcs.foundSubsequence
}
