package typeparam

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxInt8(a, b int8) int8 {
	if a < b {
		return b
	}
	return a
}

func MaxInt16(a, b int16) int16 {
	if a < b {
		return b
	}
	return a
}

func MaxInt32(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func MaxUInt(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}

// ... and the list goes on
