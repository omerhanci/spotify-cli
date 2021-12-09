package prompter

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func PromptSelect(label string, items interface{}, templates *promptui.SelectTemplates, size int ) (int, error) {
	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
		Size:      size,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}
	return i, nil
}
