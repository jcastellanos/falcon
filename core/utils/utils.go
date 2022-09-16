package utils

func CallOrDefault(call func(string) (int, error), param string, def int) int {
	res, err := call(param)
	if err != nil {
		res = def
	}
	return res
}
