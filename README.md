# Go Panic and Recover Example

This example demonstrates a common misunderstanding regarding the use of `panic` and `recover` in Go.  The code showcases a scenario where a `panic` is triggered, and the intention is to use `recover` within a `defer` function to gracefully handle the error.

**Key Point:**  `recover` only works within a `defer` function, and it must be called after the `panic` has occurred.  This example initially has the `fmt.Println("After the panic")` statement placed incorrectly, leading to unexpected behavior.

The solution provided addresses this issue by appropriately placing the code to handle the panic.