# nCAPTCHA

## How to deploy

1. Install Go.
2. Clone the repository.
3. Run `go build`.
4. Copy the binary (`ncaptcha`) and the `assets` folder to an empty directory.
5. Set the environment variable `NCAPTCHA_API` to the public URL of your ncaptcha service. This is the URL that the ncaptcha frontend will use to communicate with the backend.  
Leave this empty if you do not expect others to use your service (the demo will still work).
6. Execute the binary by running `./ncaptcha`. Then, visit `http://127.0.0.1:8080/demo`.

## Screenshots

![ncaptcha widget](https://github.com/sduoduo233/ncaptcha/blob/main/screenshots/1.png?raw=true)

![find bugs in rust code](https://github.com/sduoduo233/ncaptcha/blob/main/screenshots/2.png?raw=true)

![find wrong digits of PI](https://github.com/sduoduo233/ncaptcha/blob/main/screenshots/3.png?raw=true)

![find correct angles](https://github.com/sduoduo233/ncaptcha/blob/main/screenshots/4.png?raw=true)