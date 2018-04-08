# README

As an assumption, you've already configured `gcloud`.

Make a symbolic link to your credentials in JSON provided by GCP.
It should allow you to launch a new instance.

Then,
```
$ make apply
...

$ make ssh
...
[yourname@test-gce]$
```

Quit there once and execute the playbook.
```
$ ansible-playbook plabbook.yml
..
```

You login again, build & install.
```
$ make ssh
...

$ cd ~/go/src/github.com/tmtk75/acme-autocert-sample
$ dep init && dep ensure
```
```
$ vim main.go

$ go install
```

```
$ sudo ~/go/bin/acme-autocert-sample -d <your-domain>
```
