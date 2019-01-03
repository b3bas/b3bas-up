# B3bas Uploader

B3bas Uploader is simple S3 uploader file, build on the top of B3bas Golang Framework.

## Changelog
* Changing documentation in this [file](./CHANGELOG.md).

## Setup Environment
* Development Environment
  ```
  export B3BAS_ENV=development  
  ```
* Staging Environment
  ```
  export B3BAS_ENV=staging   
  ```
* Production Environment
  ```
  export B3BAS_ENV=production
  ```

## Setup AWS Authentication as Environment
* Setup `AWS_ACCESS_KEY`
  ```
  export AWS_ACCESS_KEY=[YOUR_AWS_ACCESS_KEY]
  ```
* Setup `AWS_SECRET_KEY`
  ```
  export AWS_SECRET_KEY=[YOUR_AWS_SECRET_KEY]
  ```
* Setup `AWS_TOKEN`
  ```
  export AWS_TOKEN=[YOUR_AWS_TOKEN]
  ```

## Setup `[environment].main.ini` File
* Define `[environment].main.ini` in `files/etc/b3bas-up/[environment]`
  ```
  [AwsConfig]
    AwsAccessKey = "YOUR_AWS_ACCESS_KEY"
    AwsSecretKey = "YOUR_AWS_SECRET_KEY"
    AwsToken = "YOUR_AWS_TOKEN"
    S3Region = "ap-southeast-1"
    S3Bucket = "b3bas-up"
  ```
* Copy your `[environment].main.ini` file to your `/etc/b3bas-up/[environment]` local machine / server
  ```
  mkdir -p /etc/b3bas-up/development
  mkdir -p /etc/b3bas-up/staging
  mkdir -p /etc/b3bas-up/production

  cp ./files/etc/b3bas-up/development/b3bas-up.main.ini /etc/b3bas-up/development/b3bas-up.main.ini
  cp ./files/etc/b3bas-up/staging/b3bas-up.main.ini /etc/b3bas-up/staging/b3bas-up.main.ini
  cp ./files/etc/b3bas-up/production/b3bas-up.main.ini /etc/b3bas-up/production/b3bas-up.main.ini
  ```  
* Default Region & Bucket in `src/config/init.go`
  ```
  S3_REGION = "ap-southeast-1"
  S3_BUCKET = "b3bas-up"
  ```

## How To
* Build (using `make`)
  ```
  make build
  ```
* Build (using `go` binary)
  ```  
  go build app.go
  ```
* Run (using `make`)
  ```
  make run
  ```
* Run (using `go` binary)
  ```
  go run app.go
  ```

## License
```
Apache version 2
```
