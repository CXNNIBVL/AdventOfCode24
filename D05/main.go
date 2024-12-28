package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	FILE      = "input"
	TEST_FILE = "input_test"
	EN_DBG    = true
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
			rules[ix].mustBeBeforePages = append(rules[ix].mustBeBeforePages, int(mustBeBeforePage))
			continue
		}

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

func applyRulesToUpdate(update []int, rules []Rule) (newUpdate []int) {
	newUpdate = make([]int, len(update))

	var applicableRules []*Rule = nil

	for _, page := range update {
		ix := slices.IndexFunc(rules, func(r Rule) bool {
			return r.page == page
		})

		if found := ix != -1; found {
			applicableRules = append(applicableRules, &rules[ix])
			fmt.Printf("Applicable: %+v\n", rules[ix])
		}
	}

	if applicableRules == nil {
		copy(newUpdate, update)
		return
	}

	// // Find last
	// for _, page := range update {
	// 	for _, rule := range applicableRules {
	// 		found := !slices.ContainsFunc(rule.mustBeBeforePages, func(e int) bool {
	// 			return slices.Contains(update, e)
	// 		})

	// 		if found {
	// 			continue
	// 		}
	// 	}
	// }

	return
}

func main() {

	file, err := os.Open(TEST_FILE)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	ruleLines, updateLines := parseRuleAndUpdateLinesFromLines(getFileLines(file))

	if ruleLines == nil || updateLines == nil {
		panic("Read no lines from file or newline separator not found")
	}

	rules := parseRules(ruleLines)

	if EN_DBG {
		fmt.Println("### Rules:")
		fmt.Println("-------------------------------------")
		sort.Slice(rules, func(i, j int) bool {
			return rules[i].page < rules[j].page
		})

		for _, rule := range rules {
			sort.Slice(rule.mustBeBeforePages, func(i, j int) bool {
				return rule.mustBeBeforePages[i] < rule.mustBeBeforePages[j]
			})
		}
	}

	for _, rule := range rules {
		fmt.Printf("%+v\n", rule)
	}

	updates := parseUpdates(updateLines)

	if EN_DBG {
		fmt.Println("### Updates:")
		fmt.Println("-------------------------------------")
		for _, update := range updates {
			fmt.Printf("%+v\n", update)
		}
	}

	// TODO: Maybe do ordering of rules first and then just walk over the list and match along the way to get the valid updates

	for _, update := range updates {
		copy(update, applyRulesToUpdate(update, rules))
	}

	if EN_DBG {
		fmt.Println("### New Updates:")
		fmt.Println("-------------------------------------")
		for _, update := range updates {
			fmt.Printf("%+v\n", update)
		}
	}
}
