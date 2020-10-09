# httpua

A Go package to set a custom User-Agent when placing HTTP requests using http.Client.

## Usage

Creating a new http client:

    client := httpua.NewClient("MyUserAgent/1.0") // Returns *http.Client
    client.Get("...")

Extending an initialized http client:

    client := httpua.WithClient(http.DefaultClient, "MyUserAgent/1.0") // Returns *http.Client
    client.Get("...")

## Documentation

https://pkg.go.dev/github.com/adamkeys/httpua
