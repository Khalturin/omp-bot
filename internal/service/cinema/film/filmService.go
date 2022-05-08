package film

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"log"
)

type filmService interface {
	Describe(filmID uint64) (*cinema.Film, error)
	List(cursor uint64, limit uint64) ([]cinema.Film, error)
	Create(cinema.Film) (uint64, error)
	Update(filmID uint64, film cinema.Film) error
	Remove(filmID uint64) (bool, error)
}

type DummyFilmService struct{}

func NewDummyFilmService() *DummyFilmService {
	return &DummyFilmService{}
}

func (fs *DummyFilmService) Describe(filmID uint64) (*cinema.Film, error) {
	return nil, nil
}

func (fs *DummyFilmService) List(cursor uint64, limit uint64) ([]cinema.Film, error) {
	if int(cursor+limit) > len(cinema.AllEntities) {
		if int(cursor) > len(cinema.AllEntities) {
			return nil, nil
		}
		return cinema.AllEntities[cursor:], nil
	}
	return cinema.AllEntities[cursor : cursor+limit], nil
}

func (fs *DummyFilmService) Create(film cinema.Film) (uint64, error) {
	cinema.AllEntities = append(cinema.AllEntities, film)
	return 0, nil
}

func (fs *DummyFilmService) Update(filmID uint64, film cinema.Film) error {
	cinema.AllEntities[filmID] = film
	return nil
}

func (fs *DummyFilmService) Remove(filmID uint64) (bool, error) {
	newAllEntities := make([]cinema.Film, 0, len(cinema.AllEntities)-1)
	j := 0
	for i, val := range cinema.AllEntities {
		if uint64(i) != filmID {
			val.ID = uint64(j)
			newAllEntities = append(newAllEntities, val)
			j++
		}

	}
	cinema.AllEntities = newAllEntities
	return true, nil
}

func (s *DummyFilmService) Get(idx int) (*cinema.Film, error) {
	if len(cinema.AllEntities) <= idx {
		log.Println("can't Get film, out of range")
		return nil, errors.New("out of range")
	}
	return &cinema.AllEntities[idx], nil
}
