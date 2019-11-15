package main

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
//type error interface {
//	Error() string
//}

type account struct {
	name string
	id string

	bins []bin
}

type bin struct {
	name string
	desc string
	cents int
}

func (acc *account) CreateBin(n string, d string, c int) {
	b := bin{
		name:  n,
		desc:  d,
		cents: c,
	}
	acc.bins = append(acc.bins, b)
}

func other() {
	//defer fmt.Println("1")
	//defer fmt.Println("2", "2")
	//var err error
	//err = errors.New("REALLY BAD ERROR")
	//
	//
	//defer fmt.Println("err: ", err)
	//defer fmt.Println("4")
}

func main() {
	//acc1 := account {
	//	name: "james and claire account",
	//	id:   "12345",
	//}
	//acc2 := account {
	//	name: "jon",
	//	id:   "2345",
	//}

	//// create a bin
	//b := CreateBin("Grocery", "food budget", 0)
	//// append the bin to account 1's bins
	//acc1.bins = append(acc1.bins, b)
	//
	//fmt.Println(b)
	//fmt.Println(acc1)
	//fmt.Println(acc2)
	// acc1.CreateBin("Grocery", "food budget", 23473)
	//acc1.CreateBin("sdf", "foosdfet", 23473)
	//acc1.CreateBin("werwer", "fsdfudget", 23473)
	//acc1.CreateBin("sdfsdf", "food asdfasdfget", 23473)
	//
	//acc2.CreateBin("SDFASDF", "fSDFSDFget", 23473)
	//acc2.CreateBin("ASDFASDFSDF", "fooSDFSDFget", 23473)
	//acc2.CreateBin("SDFSDFSDF", "fooSDFSDFSDFet", 23473)



}
