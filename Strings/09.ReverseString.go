// 	https://leetcode.com/problems/reverse-string

func reverseString(s []byte) []byte {
	byte_String := s
	var temp byte

	for i, j := 0, len(byte_String)-1; i < j; i, j = i+1, j-1 {
		temp = byte_String[i]
		byte_String[i] = byte_String[j]
		byte_String[j] = temp
	}

	return (byte_String)
}