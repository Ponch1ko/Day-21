package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Wins int    `json:"Wins"`
}

var players []Player

func readData(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка считывания данных, повторите попытку.")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Поле не должно быть пустым, попробуйте снова")
			continue
		}
		return input
	}
}

func stringToNumber(input string) (int, error) {
	val, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if val < 0 {
		return 0, fmt.Errorf("Число не может быть отрицательным, попробуйте снова!")
	}
	return val, nil
}

func menu() {
	for {
		fmt.Println("Доступные команды: ")
		fmt.Println("1. Добавить игрока")
		fmt.Println("2. Показать всех игроков")
		fmt.Println("3. Показать победителя")
		fmt.Println("4. Сохранить в файл")
		fmt.Println("5. Загрузить из файла")
		fmt.Println("6. Удалить игрока")
		fmt.Println("7. Изменить данные игрока")
		fmt.Println("8. Выход")

		command := readData("Команда: ")
		n, err := stringToNumber(command)
		if err != nil {
			fmt.Println("Ошибка при вводе, попробуйте снова!")
			continue
		}
		switch n {
		case 1:
			addPlayer()
		case 2:
			showPlayers()
		case 3:
			showWinner()
		case 4:
			saveToFile("players.json")
		case 5:
			loadFromFile("players.json")
		case 6:
			deletePlayer()
		case 7:
			changeDataForPlayer()
		case 8:
			return
		default:
			fmt.Println("Команда не найдена, попробуйте снова!")
		}
	}
}

func addPlayer() {
	name := readData("Имя нового игрока: ")
	ageStr := readData("Возраст нового игрока: ")
	age, err := stringToNumber(ageStr)
	if err != nil {
		fmt.Println("Ошибка при вводе, попробуйте снова!")
		return
	}
	winsStr := readData("Сколько побед у нового игрока: ")
	wins, err := stringToNumber(winsStr)
	if err != nil {
		fmt.Println("Ошибка при вводе, попробуйте снова!")
		return
	}
	newPlayer := Player{
		Name: name,
		Age:  age,
		Wins: wins,
	}
	players = append(players, newPlayer)
}

func showPlayers() {
	if len(players) == 0 {
		fmt.Println("Сейчас нет игроков, сначала добавьте их")
		return
	}
	for i, p := range players {
		fmt.Printf("%d. %s | %d | %d\n", i+1, p.Name, p.Age, p.Wins)
	}
}

func showWinner() {
	if len(players) == 0 {
		fmt.Println("Сейчас нет игроков, сначала добавьте их")
		return
	}
	winner := players[0]
	for _, p := range players {
		if p.Wins > winner.Wins {
			winner = p
		}
	}
	fmt.Printf("Победитель: %s | %d | %d\n", winner.Name, winner.Age, winner.Wins)
}

func saveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %s", err)
		return
	}
	defer file.Close()
	writer := json.NewEncoder(file)
	writer.SetIndent("", " ")

	if err := writer.Encode(players); err != nil {
		fmt.Println("Ошибка сохранения")
		return
	}
	fmt.Println("Данные успешно сохранены")
}

func loadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %s", err)
		return
	}
	defer file.Close()
	players = []Player{}
	reader := json.NewDecoder(file)
	if err := reader.Decode(&players); err != nil {
		fmt.Println("Ошибка декодирования")
		return
	}

	for i, p := range players {
		fmt.Printf("%d. %s | %d | %d\n", i+1, p.Name, p.Age, p.Wins)
	}
}

func deletePlayer() {
	name := readData("Введите номер игрока")
	num, err := stringToNumber(name)
	if err != nil {
		fmt.Println("Ошибка при вводе")
		return
	}
	idx := num - 1
	if idx < 0 || idx >= len(players) {
		fmt.Println("Игрок с таким номером не найден")
		return
	}

	players = append(players[:idx], players[idx+1:]...)
	fmt.Println("Игрок успешно удалён")
}

func changeDataForPlayer() {
	n := readData("Введите номер игрока: ")
	num, err := stringToNumber(n)
	if err != nil {
		fmt.Println("Ошибка при вводе")
		return
	}

	idx := num - 1
	if idx < 0 || idx >= len(players) {
		fmt.Println("Игрок с таким номером не найден")
		return
	}
	name := readData("Введите имя игрока: ")
	ageStr := readData("Введите возраст: ")
	age, err := stringToNumber(ageStr)
	if err != nil {
		fmt.Println("Ошибка при вводе")
		return
	}
	winsStr := readData("Введите количество побед: ")
	wins, err := stringToNumber(winsStr)
	if err != nil {
		fmt.Println("Ошибка при вводе")
		return
	}

	players[idx].Name = name
	players[idx].Age = age
	players[idx].Wins = wins

	fmt.Println("Данные игрока успешно обновлены")
}

func main() {
	menu()
}
