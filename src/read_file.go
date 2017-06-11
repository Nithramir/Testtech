package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"io/ioutil"
	"regexp"
)

//cette fonction vérifie la bonne taille des champs, et le format de l'email
func string_size(line []string) int {
	if len(line[0]) == 0 || len(line[0]) > 50 {
		return 1
	}
	if len(line[1]) == 0 || len(line[1]) > 50 {
		return 1
	}
	if len(line[2]) == 0 || len(line[2]) > 100 {
		return 1
	}
	err := checkmail.ValidateFormat(line[2])
	if err != nil {
		fmt.Println("Email: ", err)
		return 1
	}
	return 0
}

//Fonction verifiant le bon nombre de champs, et la bonne taille de ceux-ci
func csv_ok(file [][]string) int {
	for i := range file {
		if len(file[i]) == 4 {
			if string_size(file[i]) == 1 {
				fmt.Println("Error csv size fields.")
				return 1
			}
		} else {
			fmt.Println("Wrong format of csv file.")
			return (1)
		}
	}
	return (0)
}

//Fonction lisant le fichier passé en paremètre, et retourne son
//contenu dans une string
func read_file(name string) (string, int) {
	//verifier l'extension de fichier
	regex := regexp.MustCompile("(.+\\.csv)$")
	if regex.MatchString(name) == false {
		fmt.Println("Erreur", name, ": Wrong file Extension. Use .csv file.")
		return "", 1
	}
	data, err := ioutil.ReadFile(name)
	handle_err(err)
	return string(data), 0
}
