# gomailer

A minimal example for sending emails with gmail in go

## Instructions

+ Clone repo to `$GOPATH/src/github.com/octohedron`
+ Install dependencies

```bash
go get
```

+ Set env variables

```bash
export from_email=you@gmail.com
export email_password=hunter2
export to_email=yourotheremail@gmail.com
```

+ Build & run program

```bash
$ go build && ./gomailer &
Serving email server in port 5555
```

+ Test it with curl

```bash
$ curl -X POST \
 -F 'email=example@gmail.com' \
 -F 'name=example' \
 -F 'subject=ExampleSubject' \
 -F 'message=HelloWorld' \
  http://127.0.0.1:5555/sendemail
sent
```

Now you got the email server working, might need to add an nginx proxy, add your basic nginx configuration, then install and run `certbot` for SSL support

+ Install certbot with `apt-get install python-certbot-nginx`
+ Run certbot with `certbot --nginx`

Your new nginx.conf file should look like this

```nginx
server {
    server_name email.example.com;
    location / {
            proxy_pass http://127.0.0.1:5555;
    }
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/email.example.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/email.example.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
server {
    if ($host = email.example.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot

        listen 80;
        server_name email.example.com;
    return 404; # managed by Certbot
}
```

Reload nginx with `service nginx restart`

Now you can send emails through gmail, for example, with javascript

``` javascript
var formData = new FormData();
formData.append("name", "example name");
formData.append("email", "example email");
formData.append("subject", "example subject");
formData.append("message", "Hello World");
// make the request, change the server address
fetch("https://email.example.com/sendemail", {
  method: "POST",
  body: formData
}).then(result => {
  result.text().then(result => {
    if (result == 'sent') {
      alert("Email sent successfully :)")
    } else {
      alert("There was a problem :)")
    }
  });
});
```

