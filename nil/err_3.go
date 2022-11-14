package main

// func proverbs(name string) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	sw := safewriter{w:f}
// 	sw.writeln("A\n")
// 	sw.writeln("B\n")
// 	return sw.err
// }

// type safewriter struct {
// 	w io.Writer
// 	err error
// }

// func (sw *safewriter) writeln(s string) {
// 	if sw.err != nil {
// 		return
// 	}

// 	_, sw.err = fmt.Fprintln(sw.w, s)
// }

// func main() {

// }
