package helpers

func GetPreviousOffset(offset, limit int) int {

	previousOffset := -1

	if offset > 0 {
		countOffset := offset - limit
		if countOffset >= 0 {
			previousOffset = countOffset
		}
	}

	return previousOffset
}

func GetNextOffset(offset, limit, total int) int {
	
	nextOffset := -1

	if (offset + limit) < total {
		nextOffset = offset + limit
	}

	return nextOffset
}
