package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	FILE      = "../inputs/D05/input"
	TEST_FILE = "../inputs/D05/input_test"
)

func getFileLines(f *os.File) []string {
	lines := make([]string, 0)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func parseRuleAndUpdateLinesFromLines(lines []string) (ruleLines, updateLines []string) {
	for ix, l := range lines {
		if len(l) == 0 {
			ruleLines = lines[0:ix]
			updateLines = lines[ix+1:]
			return
		}
	}

	return nil, nil
}

type Rule struct {
	page              int
	mustBeBeforePages []int
}

func parseRules(ruleLines []string) []Rule {
	rules := make([]Rule, 0)

	for _, line := range ruleLines {
		s := strings.Split(line, "|")

		if len(s) != 2 {
			panic("Parsing Error: seperator wasnt found")
		}

		page, _ := strconv.ParseInt(s[0], 10, 64)
		mustBeBeforePage, _ := strconv.ParseInt(s[1], 10, 64)

		ix := slices.IndexFunc(rules, func(r Rule) bool {
			return r.page == int(page)
		})

		if alreadyExist := ix != -1; alreadyExist {
			// Add "before page" requirement to existing rule
			rules[ix].mustBeBeforePages = append(rules[ix].mustBeBeforePages, int(mustBeBeforePage))
			continue
		}

		// Append new rule
		rules = append(rules, Rule{page: int(page), mustBeBeforePages: make([]int, 1)})
		rules[len(rules)-1].mustBeBeforePages[0] = int(mustBeBeforePage)
	}

	return rules
}

func parseUpdates(updateLines []string) [][]int {
	updates := make([][]int, 0)

	for _, upd := range updateLines {
		s := strings.Split(upd, ",")

		update := make([]int, 0)
		for _, num := range s {
			v, _ := strconv.ParseInt(num, 10, 64)
			update = append(update, int(v))
		}
		updates = append(updates, update)
	}

	return updates
}

func applyRulesToUpdate(update []int, rules []Rule) (newUpdate []int, wasAlreadyCorrect bool) {

	wasAlreadyCorrect = true

	var applicableRules []*Rule = nil

	for _, page := range update {
		ix := slices.IndexFunc(rules, func(r Rule) bool {
			return r.page == page
		})

		if found := ix != -1; found {
			applicableRules = append(applicableRules, &rules[ix])
		}
	}

	if applicableRules == nil {
		newUpdate = append([]int{}, update...)
		return
	}

	// Insert sort
	for _, page := range update {
		newUpdate = append(newUpdate, page)
		ruleIx := slices.IndexFunc(applicableRules, func(r *Rule) bool {
			return r.page == page
		})

		// No rule found, appending is just fine
		if ruleIx == -1 {
			continue
		}

		pagerule := applicableRules[ruleIx]

		// For loop index keeps track of appended page inside array while we're swapping it to the front
		for i := len(newUpdate) - 1; i >= 1; i-- {
			// There is no requirement that the appended page needs to be in front of the one before that
			if !slices.Contains(pagerule.mustBeBeforePages, newUpdate[i-1]) {
				break
			}

			// Swap items
			wasAlreadyCorrect = false
			newUpdate[i], newUpdate[i-1] = newUpdate[i-1], newUpdate[i]
		}
	}

	return
}

func findMiddlePageNumber(update []int) int {
	if len(update)%2 == 0 {
		panic("Trying to access middle index of update, but doesn't exist on update with even length")
	}

	return update[len(update)/2]
}

func Part1And2(updates [][]int, rules []Rule) (sumCorrect, sumIncorrect int) {
	sumOfCorrectMiddlePageNrs, sumOfIncorrectMiddlePageNrs := 0, 0

	for _, update := range updates {
		newUpdate, updateWasAlreadyCorrect := applyRulesToUpdate(update, rules)
		middlePageNum := findMiddlePageNumber(newUpdate)

		if updateWasAlreadyCorrect {
			sumOfCorrectMiddlePageNrs = sumOfCorrectMiddlePageNrs + middlePageNum
		} else {
			sumOfIncorrectMiddlePageNrs = sumOfIncorrectMiddlePageNrs + middlePageNum
		}
	}

	return sumOfCorrectMiddlePageNrs, sumOfIncorrectMiddlePageNrs
}

func parseInputs() (updates [][]int, rules []Rule) {
	file, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	ruleLines, updateLines := parseRuleAndUpdateLinesFromLines(getFileLines(file))

	return parseUpdates(updateLines), parseRules(ruleLines)
}

func main() {

	updates, rules := parseInputs()
	sumCorrect, sumIncorrect := Part1And2(updates, rules)

	fmt.Println("Part1: ", sumCorrect)
	fmt.Println("Part2: ", sumIncorrect)
}
