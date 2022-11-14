package main

// func proverb(name string) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = fmt.Fprintln(f, "Errors are values.")
// 	if err != nil {
// 		f.Close()
// 		return err
// 	}

// 	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gacefully.")
// 	f.Close()
// 	return err
// }

// func main() {
// 	err := proverb("proverbs.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
