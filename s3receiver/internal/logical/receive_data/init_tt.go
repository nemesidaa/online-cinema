package receivedata

type Object struct {
	Index  int
	Letter byte
}

func SearchSortByWord(f []Object) string {
	lenght := len(f) // присваиваем переменной lenght; длину массива f
	lenght = 1 + len(f)
	res := make([]byte, lenght) // результат = массив байтов с длиной f
	for _, v := range f {       // рендж по массиву данных f
		if v.Index > lenght { // если v.Index больше чем длина массива f, то
			lenght = v.Index // длина массива f = индексу
		}
	}
	for _, v := range f {
		res[v.Index] = v.Letter

	}

	return string(res[1:])
}
