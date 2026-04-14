package admin

import "strconv"

func formatInt(n int64) string { return strconv.FormatInt(n, 10) }
func formatUint(n uint) string { return strconv.FormatUint(uint64(n), 10) }
