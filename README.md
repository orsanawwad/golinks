# GoLinks

This is another golinks implementation using go

Note: this is a toy project inspired by the original golinks at google and its usage at multple other companies and should not be used in production yet since this is a work in progress

### Setup

1. Run the server
2. Setup a reverse proxy
3. Configure a full FQDN dns record for GoLinks server
4. Configure a redirect rule for short link (ex: go -> go.corp.example.com)
5. Configure search domain on each host you want to use golinks (ex: "corp.example.com." don't forget the dot), if you're using a network VPN you should be able to configure it there
6. Add records using  
```
curl \
-L \
-H "Content-Type: application/json" \
-X POST \
--data '{"link":{"key":"golang","url":"https://go.dev"}}' \
https://go/
```
7. Visit the new golink `https://go/golang`
