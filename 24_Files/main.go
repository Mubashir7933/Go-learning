package main

func main() {
	// f, error := os.Open("textFile.txt")
	// if error != nil {
	// 	panic(error)
	// }

	// fileinfo, error := f.Stat()
	// if error != nil {
	// 	panic(error)
	// }
	// fmt.Println("Your file name is: ", fileinfo.Name())
	// fmt.Println("Your file name is: ", fileinfo.Size())

	//Reading file

	// f, err := os.Open("textFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// buf := make([]byte, 1024)
	// n, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("Number of bytes read: ", n, string(buf[i]))
	// }

	//Basic way to read file

	// f, err := os.ReadFile("textFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//holds the whole data of the file in the memory using this way

	//Reading folders

	// dir, _err := os.Open(".")
	// if _err != nil {
	// 	panic(_err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.Readdir(-1)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range fileInfo {
	// 	fmt.Println(file.Name(), file.IsDir(), file.Size(), file.Mode())
	// }

}
