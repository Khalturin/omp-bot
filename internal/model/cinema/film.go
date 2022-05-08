package cinema

import (
	"fmt"
)

var AllEntities []Film

var PageSize uint64 = 2

type Film struct {
	ID       uint64 `json:"id"`
	Year     uint64 `json:"year"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`    // Жанр
	Director string `json:"director"` // Режиссер
}

func (f *Film) String() string {
	return fmt.Sprintf(
		"ID: %d;\n Создано: %d;\n Название: %s;\n Жанр: %s;\n Режисер: %s.\n",
		f.ID,
		f.Year,
		f.Title,
		f.Genre,
		f.Director,
	)
}

func init() {
	id := uint64(0)

	AllEntities = append(AllEntities,
		Film{ID: id,
			Year:     2016,
			Title:    "Fantastic Beasts and Where to Find Them",
			Genre:    "Fantasy",
			Director: "David Yates",
		})
	id++

	AllEntities = append(AllEntities,
		Film{ID: id,
			Year:     2018,
			Title:    "Fantastic Beasts: The Crimes of Grindelwald",
			Genre:    "Fantasy",
			Director: "David Yates",
		})
	id++

	AllEntities = append(AllEntities,
		Film{ID: id,
			Year:     2022,
			Title:    "Fantastic Beasts: The Secrets of Dumbledore",
			Genre:    "Fantasy",
			Director: "David Yates",
		})
	id++

	AllEntities = append(AllEntities,
		Film{ID: id,
			Year:     2025,
			Title:    "Fantastic Beasts: 4",
			Genre:    "Fantasy",
			Director: "David Yates",
		})
	id++

	//for i := id; i < 30; i++ {
	//	AllEntities = append(AllEntities,
	//		Film{ID: i,
	//			CreatedAt: time.Date(2025, time.April, 6, 0, 0, 0, 0, time.UTC),
	//			Title:     fmt.Sprintf("Film%d", i),
	//			Genre:     "Genre",
	//			Director:  "Director",
	//			Starring:  "Actor1, Actor2, ...",
	//		})
	//}
}
