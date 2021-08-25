package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", hadleCreate)
	server.Handle("POST", "/post", handlePostRequest)
	server.Handle("POST", "/user", UserPostRequest)

	server.Listen()
}
