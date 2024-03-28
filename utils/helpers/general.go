package helpers

func GetTotalPages(totalProducts, limit int) int {
	return (totalProducts + limit - 1) / limit
}

func GetNextPage(offset, limit, totalData int) int {
	nextOffset := offset + limit
	if nextOffset < totalData {
		return (nextOffset / limit) + 1
	}
	return -1
}

func GetPrevPage(offset, limit int) int {
	if offset > 0 {
		return (offset / limit) + 1
	}
	return -1
}

func GetCurrentPage(offset, limit int) int {
	return (offset / limit) + 1
}
