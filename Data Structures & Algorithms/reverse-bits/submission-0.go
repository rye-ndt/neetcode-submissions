func reverseBits(n int) int {
	result := 0
	bits := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(bits); i++ {
		if (((n >> i) & 1) == 1) { result |= (1 << (31 - i ))}
	}
	return result
}
