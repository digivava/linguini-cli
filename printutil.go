package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Scene holds the info about a given point in the storyline
type Scene struct {
	id           string
	season       string
	day          string
	timeOfDay    string
	sceneIntro   string
	speakers     []string
	promptIntros []string
	prompts      []string
	answers      []string
}

func printScene(scene Scene, sceneSection int) {

	fmt.Printf("\n\n\n=========================\n")
	fmt.Printf("%s - %s - %s\n", strings.ToUpper(scene.season), strings.ToUpper(scene.day), strings.ToUpper(scene.timeOfDay))
	fmt.Printf("=========================\n\n")

	fmt.Printf(scene.sceneIntro)

	printPrompt(scene, sceneSection)
}

func printPrompt(scene Scene, sceneSection int) {
	fmt.Printf("\n%s\n\n", scene.promptIntros[sceneSection])

	printDialogue(scene.speakers[sceneSection], scene.prompts[sceneSection])
}

// prints "Speaker Name: Text"
func printDialogue(speaker string, text string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	fmt.Printf("%s: %s。\n\n", cyan(speaker), white(text))
}
