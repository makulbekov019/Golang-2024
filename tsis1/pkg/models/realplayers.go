package models

import "strings"

type Player struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	Number   string `json:"number"`
	Position string `json:"position"`
	Image    string `json:"image"`
}

var players = []Player{
	{Name: "Thibout", Age: "31", Number: "1", Position: "Goalkeeper", Image: "https://img.a.transfermarkt.technology/portrait/header/108390-1665067957.jpg?lm=1"},
	{Name: "Kepa", Age: "29", Number: "25", Position: "Goalkeeper", Image: "https://img.a.transfermarkt.technology/portrait/header/192279-1661855851.jpg?lm=1"},
	{Name: "Eder", Age: "26", Number: "3", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/401530-1689002515.jpg?lm=1"},
	{Name: "David", Age: "31", Number: "4", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/59016-1684921582.jpeg?lm=1"},
	{Name: "Antonio", Age: "30", Number: "22", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/86202-1684484602.jpg?lm=1"},
	{Name: "Nacho", Age: "34", Number: "6", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/86202-1684484602.jpg?lm=1"},
	{Name: "Ferland", Age: "28", Number: "23", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/291417-1701294025.jpg?lm=1"},
	{Name: "Fran", Age: "24", Number: "20", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/341264-1688119965.jpg?lm=1"},
	{Name: "Daniel", Age: "32", Number: "2", Position: "Defender", Image: "https://img.a.transfermarkt.technology/portrait/header/138927-1684232714.jpeg?lm=1"},
	{Name: "Aurelien", Age: "24", Number: "18", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/413112-1668500754.jpg?lm=1"},
	{Name: "Federico", Age: "25", Number: "15", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/369081-1681999815.jpg?lm=1"},
	{Name: "Eduardo", Age: "21", Number: "12", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/640428-1668500874.jpg?lm=1"},
	{Name: "Toni", Age: "34", Number: "8", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/31909-1700638741.jpg?lm=1"},
	{Name: "Luka", Age: "38", Number: "10", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/27992-1687776160.jpg?lm=1"},
	{Name: "Dani", Age: "27", Number: "19", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/27992-1687776160.jpg?lm=1"},
	{Name: "Jude", Age: "20", Number: "5", Position: "Midfielder", Image: "https://img.a.transfermarkt.technology/portrait/header/581678-1693987944.jpg?lm=1"},
	{Name: "Vinicius", Age: "23", Number: "7", Position: "Forward", Image: "https://img.a.transfermarkt.technology/portrait/header/371998-1664869583.jpg?lm=1"},
	{Name: "Rodrygo", Age: "23", Number: "11", Position: "Forward", Image: "https://img.a.transfermarkt.technology/portrait/header/371998-1664869583.jpg?lm=1"},
	{Name: "Alvaro", Age: "19", Number: "29", Position: "Forward", Image: "https://img.a.transfermarkt.technology/portrait/header/948275-1660124773.jpg?lm=1"},
	{Name: "Joselu", Age: "33", Number: "14", Position: "Forward", Image: "https://img.a.transfermarkt.technology/portrait/header/81999-1686816784.jpg?lm=1"},
	{Name: "Karrilo", Age: "21", Number: "34", Position: "Forward", Image: "https://cdn-m.sport24.ru/m/5869/e001/a3b2/41f2/ad91/f814/1015/60c1/160_10000_max.png"},
}

func GetAllPlayers() []Player {
	return players
}

func GetPlayerByName(name string) Player {
	for i := 0; i < len(players); i++ {
		player := players[i]
		nameWithoutSpaces := strings.ReplaceAll(player.Name, " ", "")
		nameWithoutSpaces = strings.ReplaceAll(nameWithoutSpaces, "/", "")
		if nameWithoutSpaces == name {
			return player
		}
	}
	return Player{}
}
