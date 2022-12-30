# cloudflare

A provider for (initially) basic interactions with cloudflare DNS to support partners that are using cloudflare.

Example:

```golang
    p, err := cloudflare.NewProviderFromConfig(&cloudflare.Config{
        // the name of a DNS zone in
        // cloudflare, the Select process
        // will scan and find the appropriate ID
		Zone:  "wutangclan.com",

        // the name of the A record to create
		Name:  "gza",

        // the IP value
		Value: "121.21.2.1",

        // (optional) TTL - will default to 1 (auto)
        // if left empty
		TTL:   1,
	})

	if err != nil {

		return err
	}

	err = p.Apply()
	if err != nil {
		return err
	}
```
