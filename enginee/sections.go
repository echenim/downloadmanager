package enginee

// if file size is 100 bytes, our section should like:
// [[0 10] [11 21] [22 32] [33 43] [44 54] [55 65] [66 76] [77 87] [88 98] [99 99]]
func (d Downloader) FormSections(sections [][2]int, eachSize int) [][2]int {
	for i := range sections {
		if i == 0 {
			//starting byte of first section
			sections[1][0] = 0
		} else {
			//starting byte of other section
			sections[i][0] = sections[i-1][1] + 1
		}

		if i < d.Section-1 {
			sections[i][1] = sections[i][0] + eachSize
		}

	}

	return sections
}
