package entities

// Route is a struct that holds an entrypoint and multiple destinations.
// Its role is to store a link for our reverse proxy. An entrypoint is the url the client type that will send them to the reverse proxy,
// the destinations are the backend server that the reverse proxy will redirect the request to.
type Route struct {
	Entrypoint   string
	Destinations []string

	// We could add some attributes to overrides the header etc...
	// but let's keep it simple for now
}
