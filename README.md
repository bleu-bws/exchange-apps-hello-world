# exchange-apps-hello-world

## create app
```
{
    "title": "Example App",
    "description": "App com backend",
    "name": "Example App",
    "slug": "exampleapp",
    "createdBy": "example.com",
    "uri": "https://example.com",
    "scope": ["profile"],
    "redirect_uris": ["https://example.com"],
    "kind": "app",    
    "image": "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"
}
```

## get client credentials

```
http://exchange.com/api/v3/private/apps/exampleapp/credentials
```