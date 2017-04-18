/*Package reservoire ...
 * container for reservoire sampling
 * Fabian Peters 2017
 */
package reservoire

import (
	"errors"
	"math/rand"
)

//StringReservoire : reservoire container for strings
type StringReservoire struct {
	counter int
	items   []string
	rnd     rand.Rand
}

//Add adds a new element to the reservoire
func (r *StringReservoire) Add(s string) {
	if r.counter < len(r.items) {
		// if the reservoire isn't full, keep filling it
		r.items[r.counter] = s
	} else {
		// otherwise find out whether to replace an existing item
		ind := r.rnd.Intn(r.counter + 1)
		if ind < len(r.items) {
			r.items[ind] = s
		}
	}
	r.counter++
}

//GetAll returns an array containing all stored items
func (r *StringReservoire) GetAll() []string {
	l := r.Len()
	return r.items[:l]
}

//Len returns the number of items in the reservoire
func (r *StringReservoire) Len() int {
	capacity := len(r.items)
	if r.counter < capacity {
		return r.counter
	}
	return capacity
}

//NewStringReservoire creates an empty StringReservoire instance
func NewStringReservoire(capacity int, seed int64) (StringReservoire, error) {
	if capacity < 1 {
		return StringReservoire{}, errors.New("Invalid capacity.")
	}

	rnd := rand.New(rand.NewSource(seed))
	strArr := make([]string, capacity)

	return StringReservoire{items: strArr, rnd: *rnd}, nil
}
