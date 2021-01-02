package main

import (
	"fmt";
	"strings";
	"os";
	"os/exec";
	"io/ioutil";
	"encoding/json"
)

var line string = "=-=-=-=-=-=-=-=-=-=-=-="

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func newNote() {
	var fName string
	fmt.Print("Nome da nova nota: ")
	fmt.Scan(&fName)

	err := ioutil.WriteFile(("./notes/" + fName + ".txt"), []byte(""), 0755)
	check(err)

	setRec(fName)
}

func showNotes() {
	files, err := ioutil.ReadDir("./notes")
	check(err)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func readNote() {
	var note string

	fmt.Print(line)
	fmt.Println()
	showNotes()
	fmt.Println(line)

	fmt.Print("Escolha uma nota para ler: ")
	fmt.Scanln(&note)

	txt, err := ioutil.ReadFile("./notes/" + note + ".txt")
	check(err)

	fmt.Println()
	fmt.Println(string(txt))
	fmt.Println()

	fmt.Scanln()

	setRec(note)
}

func editNote() {
	var note string

	fmt.Print(line)
	fmt.Println()
	showNotes()
	fmt.Println(line)

	fmt.Print("Escolha uma nota: ")
	fmt.Scanln(&note)

	cmd := exec.Command("cmd", "/c", "notepad", "notes/" + note + ".txt")
    cmd.Stdout = os.Stdout
	cmd.Run()
	
	setRec(note)
}

func deleteNote() {
	var note string

	fmt.Print(line)
	fmt.Println()
	showNotes()
	fmt.Println(line)

	fmt.Print("Escolha uma nota para excluir: ")
	fmt.Scanln(&note)

	cmd := exec.Command("cmd", "/c", "del", "notes\\" + note + ".txt")
    cmd.Stdout = os.Stdout
	cmd.Run()
}

func getRec() []string {
	jsonFile, err := os.Open("rec.json")
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var notes []string
	json.Unmarshal(byteValue, &notes)

	return notes
}

func setRec(note string) {
	notes := getRec()

	notes[4] = notes[3]
	notes[3] = notes[2]
	notes[2] = notes[1]
	notes[1] = notes[0]
	notes[0] = note

	jsonFile, _ := json.MarshalIndent(notes, "", " ")
	err := ioutil.WriteFile("rec.json", jsonFile, 0644)
	check(err)
}

func main () {
	for {
		var op string

		fmt.Println("=-=-= MNEMON =-=-=")
		fmt.Println("")
		fmt.Println("Arquivos recentemente abertos")
		for _, note := range getRec() {
			fmt.Println(note)
		}
		fmt.Println("")
		fmt.Println("Bem-vindo(a), o que gostaria de fazer?")
		fmt.Print("> ")
		fmt.Scanln(&op)

		op = strings.ToLower(op)

		switch op {
			case "n":
				newNote()
			case "r":
				readNote()
			case "s":
				showNotes()
				fmt.Scanln()
			case "rn":
				editNote()
			case "e":
				editNote()
			case "d":
				deleteNote()
			case "q":
				return
			default:
				fmt.Println("Desculpe, essa não é uma opção válida")
		}
		cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}