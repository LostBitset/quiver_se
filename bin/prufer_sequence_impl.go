package main

func BakePruferSequence(sequence []int) (ps PruferSequence) {
	ps.sequence = sequence
	n := len(sequence)
	occurences := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		occurences[i] = 1
	}
	for _, item := range sequence {
		occurences[item] += 1
	}
	ps.occurences = occurences
	return
}

func (ps PruferSequence) ToTree() (tree SimpleTree) {

}
