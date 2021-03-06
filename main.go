package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CrowdSurge/banner"
)

// Scene holds the info about a given point in the storyline
type Scene struct {
	id         string
	season     string
	day        string
	timeOfDay  string
	sceneIntro string
	prompts    []Prompt
}

// Prompt does prompty things val will fill out later
type Prompt struct {
	speaker string
	intro   string
	text    string
	answer  string
}

var scenes = []Scene{
	{
		id:        "001",
		season:    "summer",
		day:       "Monday",
		timeOfDay: "morning",
		sceneIntro: fmt.Sprintf("%-6s\n\n%-6s\n\n%-6s\n\n%-6s\n", `"Cock-a-doodle-doo", buzzes the chainsaw--or wait, the cicada--outside your window. Welcome to yet another hot and sticky death-march in the town of Fujimino, Saitama prefecture. Your exchange student orientation really didn't paint a realistic expectation of how hot summers are in Japan, or how hard it would be to find deodorant here.`,
			`"I came here for life-changing sociocultural insights and Pikachu", you remind yourself.`,
			`You mash the button on the remote for the AC unit. It produces a soothing beep-beep to assure you of future cold air. After a few minutes of scrolling through your dual-language social media feeds (melodramatic debating on the English feed, a what's-it-gonna-be-this-time roulette of cute parfaits and explicit gay cartoon porn on the Japanese), you roll out of bed and wash your face in the bathroom. You showered last night, but it's time to shower again.`,
			`Once presentable-ish, you don your school uniform--the short-sleeved sailor-top kind, just like your weeby dreams prepared you for--and head downstairs.`),
		prompts: []Prompt{
			{
				speaker: "ゆうた",
				intro:   fmt.Sprintf("Your older host brother Yuuta, 18 years old, is in the kitchen, yawning and buttering some toast. He notices you and greets you."),
				text:    "おはよう",
				answer:  "おはよう",
			},
			{
				speaker: "ゆうた",
				intro:   "He hands you some of his toast.",
				text:    "どうぞ",
				answer:  "ありがとう",
			},
		},
	},
	{
		id:         "002",
		season:     "season2",
		day:        "day2",
		timeOfDay:  "timeOfDay2",
		sceneIntro: "sceneIntro2",
		prompts: []Prompt{
			{
				speaker: "Billy Bob Thornton Place",
				intro:   fmt.Sprintf("whee"),
				text:    "howdy",
				answer:  "hidee ho",
			},
			{
				speaker: "The Duke",
				intro:   "A stranger hands you some of his smores.",
				text:    "here",
				answer:  "heck yes",
			},
		},
	},
	// {id: "002", season: "summer", day: "Monday", timeOfDay: "morning", sceneIntro: "This is the next scene yay", speakers: []string{"おかあさん"}, promptIntros: []string{"prompt1yay", "prompt2yay", "prompt3yay"}, prompts: []string{"sent1yay", "sent2yay"}},
}

func main() {

	banner.Print("linguini")

	for _, scene := range scenes {

		playScene(scene)
	}

}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	userInput, _ := reader.ReadString('\n')
	// // convert CRLF to LF
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func compareInput(userInput string, correctAnswer string) {
	if strings.Compare(correctAnswer, userInput) == 0 {
		fmt.Println("You got it right!")
	} else {
		fmt.Println("Try again!")
	}
}
