// continue에 레이블 사용

x := 7
table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
for row := 0; row < len(table); row++ {
	for col := 0; < len(table[row]); col++ {
		if table[row][col] == x {
			fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
			// LOOP로 지정된 for 문의 다음 반복문 수행
			continue LOOP
		}

	}
}
