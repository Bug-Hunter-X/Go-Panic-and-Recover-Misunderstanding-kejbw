func main() {



        fmt.Println("Before the panic")

        defer func() {

                if r := recover(); r != nil {

                        fmt.Println("Recovered from panic:", r)

                }

        }()

        panic("This is a panic")

        //Corrected position

        //fmt.Println("After the panic") //This line should not execute after a panic

} 