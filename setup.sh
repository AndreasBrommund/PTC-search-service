echo 'go get gorilla/context'
go get github.com/gorilla/context

echo 'go get julienschmidt/httprouter'
go get github.com/julienschmidt/httprouter

echo 'go get justinas/alice'
go get github.com/justinas/alice

echo 'go get justinas/lib/pq'
go get github.com/lib/pq

echo 'go get gopkg.in/olivere/elastic.v3'
go get gopkg.in/olivere/elastic.v3

echo 'Add config folder'
mkdir config

echo 'Add db config file'
echo '{
	"dev": {
		"database": "postgres",
		"user": "",
		"password": "",
		"host": "127.0.0.1",
		"port": "5432",
		"ssl_mode": "disable"
	}
}' > config/db.json
