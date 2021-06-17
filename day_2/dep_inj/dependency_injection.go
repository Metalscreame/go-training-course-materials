package dep_inj

/*
In software engineering, dependency injection is a technique in which an object receives other objects that it depends on.
These other objects are called dependencies.
 */

/*
Wrong

type Server struct {
	repository Repository
}

func New() *Server {
	return &Server{
		repository: buildMyRepo(),
	}
}


OK
type Server struct {
  repository Repository
}

func New(r *Repository) *Server {
  return &Server{
    repository: r
  }
}
 */
