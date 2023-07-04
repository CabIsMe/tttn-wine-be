## Installation & Start
1. In the source project directory run this command to get dependency library
```
go mod download
```
or try this command
```
go get
 ```
2. Go mod tidy cleans up unused dependencies or adds missing dependencies
```
go mod tidy
```
3. Make sure in your GOPATH has those install library. Something like this\
    I has $GOPATH locate in this path
```
vmd@vmd:echo $GOPATH
/home/vmd/go
ls $GOPATH/pkg/mod/github.com
 ajg             coreos        go-check        go-stack          karrick         mergermarket   rogpeppe        stretchr
 andybalholm     cosiner       go-delve        grpc-ecosystem    kataras         microcosm-cc   russross        subosito
 aquasecurity    cpuguy83      gofiber         hashicorp         kkyr            mitchellh      ryanuber        teramoby
 armon           cweill        gofrs           haya14busa        klauspost       modern-go      satori          tmc
 aymerick        davecgh       gogo            imkira           '!knetic'        moul           schollz         twinj
 becheran        denisenkom    golang          inconshreveable   konsorten       natefinch      sergi           ugorji
 benbjohnson     dgrijalva     golang-jwt      iris-contrib      kr              nsf            shareed2k       uudashr
 beorn7          dgryski       golang-sql      jackal-xmpp       labstack        onokonem      '!shopify'       valyala
 bgentry         eknkc         go-mysql-org    jackc             lann            ortuman
``` 
5. If you dont have "make" install you could run the following cmd. Step by step\
    Go will create executable file by run this cmd.
```
go build -o fc_optimal_assignment
```
6.  
	Make sure you have ".env" file locate in root directory then run this cmd.
```
./fc_optimal_assignment
```
7.	Your env file must have these key-value or in variable environment

| Key                  |               Description               |
| :--------------------| -------------------------------------:  |
| USE_PRODUCTION       |                 0 or 1                  |
| DB_HOST              |             Host of the DB              |
| DB_USERNAME          |               DB username               |
| DB_PASSWORD          |               DB password               |
| DB_DATABASE          |           Name of DB instance           |
| DB_PORT              |           Port connect to DB            |
