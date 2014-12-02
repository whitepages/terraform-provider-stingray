# go-stingray [![Build Status](https://secure.travis-ci.org/whitepages/go-stingray.png)](http://travis-ci.org/whitepages/go-stingray)

go-stingray is a client library for accessing the Stingray Traffic
Manager REST API.

For documentation of the Stingray API, see the "Stingray Traffic
Manager REST API Guide" version 2.6, March 2014 from the
[Riverbed support site](https://support.riverbed.com/content/support/software/steelapp/traffic-manager.html).

## Usage

```go
import "github.com/whitepages/go-stingray"
```

Create a new Stingray client.

```go
client := stingray.NewClient(nil, "https://localhost:9070", "username", "password")
```

Manage an extra file.

```go
fmt.Println("Writing...")
r := stingray.NewExtraFile("name")
r.Content = []byte("Test")
resp, err := client.Set(r)
fmt.Println("Status:", resp.Status)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Reading...")
r, resp, err = client.GetExtraFile("name")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Status:", resp.Status)
fmt.Println("Content:", string(r.Content))

fmt.Println("Deleting...")
r = stingray.NewExtraFile("name")
resp, err = client.Delete(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println("Status:", resp.Status)
```

```sh
Writing...
Status: 201 Created
Reading...
Status: 200 OK
Content: Test
Deleting...
Status: 204 No Content
```

## Supported Resources

Support for resources is being added as needed.

- [x] Action Program
- [ ] Alerting Action
- [ ] Aptimizer Application Scope
- [ ] Aptimizer Profile
- [ ] Bandwidth Class
- [ ] Cloud Credentials
- [ ] Custom configuration set
- [ ] Event Type
- [x] Extra File
- [ ] GLB Service
- [ ] Global Settings
- [x] License
- [ ] Location
- [x] Monitor
- [x] Monitor Program
- [ ] NAT Configuration
- [x] Pool
- [ ] Protection Class
- [x] Rate Shaping Class
- [x] Rule
- [x] SLM Class
- [ ] SSL Client Key Pair
- [x] SSL Key Pair
- [x] SSL Trusted Certificate
- [ ] Security Settings
- [ ] Session Persistence Class
- [x] Traffic IP Group
- [x] Traffic Manager (incomplete)
- [ ] TrafficScript Authenticator
- [ ] User Authenticator
- [ ] User Group
- [x] Virtual Server
