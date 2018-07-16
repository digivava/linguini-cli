package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func printSpeakerName(speaker string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	speakerNameFormatted := fmt.Sprintf("%s", cyan(speaker))
	return speakerNameFormatted
}

func playScene(scene Scene) {

	printSetting(scene)
	printSceneIntro(scene.sceneIntro)

	for _, prompt := range scene.prompts {
		printPrompt(prompt)
		in := getInput()
		compareInput(in, prompt.answer)
	}

}

func printSceneIntro(sceneIntro string) {
	fmt.Printf(sceneIntro)
}

func printSetting(scene Scene) {
	d := color.New(color.FgMagenta, color.Bold)

	fmt.Printf("\n\n\n=========================\n")
	d.Printf("%s - %s - %s\n", strings.ToUpper(scene.season), strings.ToUpper(scene.day), strings.ToUpper(scene.timeOfDay))
	fmt.Printf("=========================\n\n")
}

func printPrompt(prompt Prompt) {

	printPromptIntro(prompt.intro)
	printPromptText(prompt.speaker, prompt.text)
}

func printPromptIntro(intro string) {
	magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
	fmt.Printf("\n%s\n\n", magenta(intro))
}

// prints "Speaker Name: Text"
func printPromptText(speaker string, text string) {
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	fmt.Printf("%v: %s。\n\n", printSpeakerName(speaker), white(text))
}
