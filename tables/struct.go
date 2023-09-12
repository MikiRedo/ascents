package tables

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
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`