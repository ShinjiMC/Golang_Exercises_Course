package main

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

type ElectionResult struct {
	Name  string // Name of the candidate
	Votes int    // Votes of votes the candidate had
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var T ElectionResult
	T.Name = candidateName
	T.Votes = votes
	return &T
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}

func main() {
	initialVotes := 2
	counter := NewVoteCounter(initialVotes)
	fmt.Println("Test 1: ", *counter == initialVotes)
	fmt.Println("Test 2: ", VoteCount(counter))
	IncrementVoteCount(counter, 2)
	fmt.Println("Test 3: ", VoteCount(counter), " - ", initialVotes)
	result := NewElectionResult("Peter", 3)
	fmt.Println("Test 4: ", result.Name, " - ", result.Votes)
	fmt.Println("Test 5: ", DisplayResult(result))
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}
	DecrementVotesOfCandidate(finalResults, "Mary")
	fmt.Println("Test 6: ", finalResults["Mary"])
}
