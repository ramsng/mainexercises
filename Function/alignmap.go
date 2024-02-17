package alignmap()

func alignmap(mapi map[string]string)map[string]string {
	var maxsize int
	//identify maximum length of the map input
	for a, _ := range mapi {
		if maxsize < len(a) {
			maxsize = len(a)
		}
	}
	//equalize the length with maximum length variable for printing
	for i, j := range mapi {
		if size := maxsize - len(i); size > 0 {
			for k := 0; k < size; k++ {
				i = i + " "
			}
		}
		fmt.Printf(" Person : %v\t\tLikes : \t\t%v\n", i,j)
	}
}
