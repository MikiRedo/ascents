package tables

import (
	"gorm.io/gorm"
	"fmt"
)


type Ascents struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Grade string		`json:"grade"`
	Tries string		`json:"tries"`
	Date string			`json:"date"`
	Crag string			`json:"crag"`
	Area string			`json:"area"`
	Obs string			`json:"obs"`
} 

const AscentsSchema string = `CREATE TABLE Ascents (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR (50) NOT NULL,
	grade VARCHAR (50) NOT NULL,
	tries VARCHAR (50) NOT NULL,
	date DATE,
	crag VARCHAR (50),
	area VARCHAR (50),
	obs VARCHAR (50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

func Intersect(a, b []Ascents) []Ascents {
    set := make(map[Ascents]bool)
    var result []Ascents

    for _, line := range a {
        set[line] = true
    }

    for _, line := range b {
        if set[line] {
            result = append(result, line)
        }
    }

    return result
}

func FiltrarGrau(db *gorm.DB, grade string) ([]Ascents, error) {
	var lines []Ascents

	result := db.Where("grade=?", grade).Find(&lines) 

	if result.Error != nil {
		fmt.Printf("Nothing here")
	}
	return lines, nil					//return all the lines from the DB with value grade="?"
}

func FilterArea(db *gorm.DB, area string) ([]Ascents, error) {
	var lines []Ascents

	result := db.Where("area=?", area).Find(&lines)

	if result.Error != nil {
		fmt.Printf("Nothing in this area")
	}
	return lines, nil					//return all the lines from the DB with value area="?"
}