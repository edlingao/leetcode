package problems
func IsOneBitCharacter(bits []int) bool {
	endsInA := false
	for i := 0; i < len(bits); i++ {
		window := i + 1

		if window > len(bits)-1 {
			endsInA = true
		}

		if bits[i] == 1 && (bits[window] == 0 || bits[window] == 1) {
			i += 1
		}
	}

	return endsInA
}
