# Chi-rest

Boilerplate for REST service using go-chi. This project uses the concept of database agnostic user models by implementing storer interface for the model. 

## TODO
* improve errors in validation

## Dependencies

	* github.com/oxequa/realize
	* github.com/go-chi/chi
	* github.com/go-chi/render
	* github.com/go-chi/middleware
	* github.com/go-chi/jwtauth
	* github.com/sirupsen/logrus
	* github.com/unrolled/render
	* github.com/josharian/impl - interface implementation generator
	* github.com/jinzhu/gorm - ORM
	* github.com/go-gormigrate/gormigrate - database migration
	* github.com/go-ozzo/ozzo-validation - struct validation. why this and not other validation package that use struct tags? i really like the concept of this validation package. it doesn't use struct tags and that's what makes it (IMO) superior to others!

## Running
`go get` all dependencies before anything else then run:
    
    $ realize start

## Tools that might be needed for convenience

* Model generator for a given database
* Response and Request management package
* Stub generator for resources
* Storer interface implemetation generators - impl
* Store implementation tests stub generator
* Database migration - gormigrate (inside store implementations)
* ORM - gorm
	
### Generating Storer interface stub
Below is an example of generating Storer interface sub. 

	$ impl "s *Store" github.com/sf9v/harper/model/user.Storer >> ./store/user/mssql/mssql.go
