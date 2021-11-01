package main

func checkInclusion(s1 string, s2 string) bool {


	res := []string{}
	path := ""

	pailie(0, s1, path, &res)




}

func pailie(startindex int, s1 string, path string, res *[]string){

	if len(path) == len(s1){
		tmp := make([]byte, len(path))
		copy(path, tmp)
		*res = append(res, tmp)
		return
	}

	for i:=startindex;i<len(s1);i++{
		path = append(path, s1[i])
		pailie(i+1; s1, path, res)
		path = path[:len(path)-1)]
}


}