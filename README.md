# Harper 

Boilerplate for REST service using go-chi. This project uses the concept of database agnostic user models by implementing storer interface for the model. 

## Dependencies

	* github.com/pilu/fresh
	* github.com/go-chi/chi
	* github.com/go-chi/render
	* github.com/go-chi/middleware
	* github.com/go-chi/jwtauth
	* github.com/sirupsen/logrus
	* github.com/unrolled/render
	* github.com/josharian/impl - interface implementation generator

## Running
`go get` all dependencies before anything else then run:
    
    $ fresh

## Tools that might be needed for convenience

* Model generator for a given database
* Response and Request management package
* Storer interface implemetation generators
	
### Generating Storer interface stub
Below is an example of generating Storer interface sub. 

	$ impl "s *Store" github.com/moqafi/harper/model/user.Storer >> ./store/user/mssql/mssql.go
