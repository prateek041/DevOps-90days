package main

import (
	"fmt"
	"log"
	"net/http"
)

// func testFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "this was testing the connection")
// }

func helloHandler(w http.ResponseWriter, r *http.Request) { // handler function has two args, w for response writing and r pointer for request.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello, I am the hello handler function")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Form POST request was succesful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // tells the base URL from where the server will starting handling requests.

	http.Handle("/", fileServer)          // Registers the handler.
	http.HandleFunc("/form", formHandler) // Register the handler function for a given pattern.
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting starting at port 8080\n")
	err := http.ListenAndServe(":8080", nil) // http.HandlerFunc is used to convert a normal function {testFunc} into a HTTP handler
	if err != nil {
		log.Fatal(err)
	}
}

// What I have learnt.

// 1: Every file in go must begin with a package name.
// 2: then we import some packages that contain the functions needed for various tasks like printing to handling http requests.
// 3: main, is from where the execution of the program begins.
// 4: FileServer is used to serve files over HTTP from a specific directory.
// 5: http.Dir(root) tells the root directory (base url) and perform further actions.

// 6: http.ListenAndServe('address string', handler http.Handler) error.
//		This function accepts two arguments, {address string} which will tell which port to listen to, and handler function of type http.Handler, which can be nil.
// 		The test func tells what to run.
//		It returns a non-nil error if generated.

// 7: http.Handle(Pattern String, handler) tells what to do with the specific route. Will look more into it.
// 8: http.HandleFunc(Pattern String, handler func(ResponsWriter, *request))

// 9: http.Error(w ResponseWriter, error String, code int).
//		writes the {error string} and {error code} into the ResponseWriter type value, w and ends the request. in short, replies to the request with error string/code.

// 10: fmt.Fprintf( destination, format String) (count int, err error).
//		writes the {format String} into the destination.
//		returns: count is the number of bytes written and error is the error, if any.

// 11: func formHandler (w http.ResponseWriter, r *http.Request) this function is used as an HTTP handler.
//		[Type] ResponseWriter: is an interface that is used by the handler to write an HTTP response.
//		[Type] Request: is a struct and contains a lots of information about the request recieved. information like request method, URL etc.
//		w and r are the values of type ResponseWriter and Request. they can be named anything.
// 12:
