## Stripped down version of [gohttpserver](https://github.com/codeskyblue/gohttpserver) for internal use.

These features was removed:
 - Support show or hide hidden files
 - All authentication methods 
 - Plist proxy

## Usage
Listen on port 8000 of all interfaces, and enable file uploading.

```
$ gohttpserver -r ./ --port 8000 --upload
```

Use command `gohttpserver --help` to see more usage.

- Enable upload

  ```sh
  $ gohttpserver --upload
  ```

- Enable delete and Create folder

  ```sh
  $ gohttpserver --delete
  ```

## Advanced usage
Add access rule by creating a `.ghs.yml` file under a sub-directory. An example:

```yaml
---
upload: false
delete: false
users:
- tokens:
   - 4567gf8asydhf293r23r
  delete: true
  upload: true
```

`tokens` is a list of tokens used for upload. see [upload with curl](#upload-with-curl)

Token can be set by url `?token=4567gf8asydhf293r23r` or by cookie.
You can set `token` cookie by `http://localhost:8000/-/login/4567gf8asydhf293r23r` request.

For example, in the following directory hierarchy, users can delete/uploade files in directory `foo`, but he/she cannot do this in directory `bar`.

```
root -
  |-- foo
  |    |-- .ghs.yml
  |    `-- world.txt 
  `-- bar
       `-- hello.txt
```

User can specify config file name with `--conf`, see [example config.yml](testdata/config.yml).

To specify which files is hidden and which file is visible, add the following lines to `.ghs.yml`

```yaml
accessTables:
- regex: block.file
  allow: false
- regex: visual.file
  allow: true
```

### Upload with CURL
For example, upload a file named `foo.txt` to directory `somedir`

```sh
$ curl -F file=@foo.txt localhost:8000/somedir
{"destination":"somedir/foo.txt","success":true}

# upload with token
$ curl -F file=@foo.txt -F token=12312jlkjafs localhost:8000/somedir
{"destination":"somedir/foo.txt","success":true}

# upload and change filename
$ curl -F file=@foo.txt -F filename=hi.txt localhost:8000/somedir
{"destination":"somedir/hi.txt","success":true}
```

Note: `\/:*<>|` are not allowed in filenames.

### Deploy with nginx
Recommended configuration, assume your gohttpserver listening on `127.0.0.1:8200`

```
server {
  listen 80;
  server_name your-domain-name.com;

  location / {
    proxy_pass http://127.0.0.1:8200; # here need to change
    proxy_redirect off;
    proxy_set_header  Host    $host;
    proxy_set_header  X-Real-IP $remote_addr;
    proxy_set_header  X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header  X-Forwarded-Proto $scheme;

    client_max_body_size 0; # disable upload limit
  }
}
```

gohttpserver should started with `--xheaders` argument when behide nginx.

Refs: <http://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size>

gohttpserver also support `--prefix` flag which will help to when meet `/` is occupied by other service. relative issue <https://github.com/codeskyblue/gohttpserver/issues/105>

Usage example:

```bash
# for gohttpserver
$ gohttpserver --prefix /foo --addr :8200 --xheaders
```

**Nginx settigns**

```
server {
  listen 80;
  server_name your-domain-name.com;

  location /foo {
    proxy_pass http://127.0.0.1:8200; # here need to change
    proxy_redirect off;
    proxy_set_header  Host    $host;
    proxy_set_header  X-Real-IP $remote_addr;
    proxy_set_header  X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header  X-Forwarded-Proto $scheme;

    client_max_body_size 0; # disable upload limit
  }
}
```

## FAQ

### How the query is formated
The search query follows common format rules just like Google. Keywords are seperated with space(s), keywords with prefix `-` will be excluded in search results.

1. `hello world` means must contains `hello` and `world`
1. `hello -world` means must contains `hello` but not contains `world`

## Developer Guide
Depdencies are managed by [govendor](https://github.com/kardianos/govendor)

1. Build develop version. **assets** directory must exists

  ```sh
  $ go build
  $ ./gohttpserver
  ```
2. Build single binary release

  ```sh
  $ go build
  ```

Theme are defined in [assets/themes](assets/themes) directory. Now only two themes are available, "black" and "green".


## Reference Web sites

* Core lib Vue <https://vuejs.org.cn/>
* Icon from <http://www.easyicon.net/558394-file_explorer_icon.html>
* Code Highlight <https://craig.is/making/rainbows>
* Markdown Parser <https://github.com/showdownjs/showdown>
* Markdown CSS <https://github.com/sindresorhus/github-markdown-css>
* Upload support <http://www.dropzonejs.com/>
* ScrollUp <https://markgoodyear.com/2013/01/scrollup-jquery-plugin/>
* Clipboard <https://clipboardjs.com/>
* Underscore <http://underscorejs.org/>

**Go Libraries**

* [vfsgen](https://github.com/shurcooL/vfsgen) Not using now
* [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs) Not using now
* <http://www.gorillatoolkit.org/pkg/handlers>

## History
The old version is hosted at <https://github.com/codeskyblue/gohttp>

## LICENSE
This project is licensed under [MIT](LICENSE).
