// Interface:
// Good for Unit Tests. You just create a new struct to do the mock function,
// Then implement that interface to do the unit testing. (followed by dependency inversion of SOLID.)
package repositories

import "github.com/DagmarC/clean-arch-example/cockroach/entities"

type Cockroach interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
