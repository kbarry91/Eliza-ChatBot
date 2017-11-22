package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//create struct of responses
type Response struct {
	Patterns *regexp.Regexp
	Answers  []string
}

/*
*	function to make responses and populate struct
*	Read file line by line to slice
*	split each string using ;
*	make patterns case insensitive
* 	append values to struct
 */
func makeResponses(path string) []Response {
	fullFile, _ := ReadLines(path)
	responses := make([]Response, 0)
	for i := 0; i < len(fullFile); i += 2 {
		allPatterns := strings.Split(fullFile[i], ";")
		allResponses := strings.Split(fullFile[i+1], ";")
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern
			Patterns := regexp.MustCompile(pattern)
			responses = append(responses, Response{Patterns: Patterns, Answers: allResponses})
		}
	}
	return responses
}

// function to test if responses are being populated
func PrintResponses(path string) {
	response := makeResponses(path)
	fmt.Printf("%+v\n", response)
}

// function to read all lines from file
// adapted from https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readLine := scanner.Text()

		if skipComment(readLine) {
			continue
		}
		lines = append(lines, scanner.Text())
		//Response.Pattern = scanner.Text()
	}
	return lines, scanner.Err()
}

// if file line begins with
func skipComment(readLine string) bool {
	return strings.HasPrefix(readLine, "//") || len(strings.TrimSpace(readLine)) == 0
}

/*function to map prounouns
* 	splits string into slice of strings
*	check if any strings match the map strings
*	set string to map value
*	join up into one string again and return
 */
func matchPronouns(inputStr string) string {
	// split inputStr into slice of strings
	splitStr := strings.Fields(inputStr)

	//map of reflected pronouns
	pronouns := map[string]string{
		"i":      "you",
		"was":    "were",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "I",
		"me":     "you",
		"me.":    "you",
		"you're": "I’m",
	}

	//loop map and check for simularites
	// swap values
	for index, word := range splitStr {
		if value, ok := pronouns[strings.ToLower(word)]; ok {
			splitStr[index] = value
		}
	}
	return strings.Join(splitStr, " ")
}

/*
* 	function to find input word so it can be returned to user
*	eg. "my name is .." > "your name is"
*   resource used https://golang.org/pkg/regexp/
 */
func wordSwapper(pattern *regexp.Regexp, input string) string {
	match := pattern.FindStringSubmatch(input)
	if len(match) == 1 {
		return "" // no wordmatch is required
	}
	wordSwap := match[1]               // 1 is first match.
	wordSwap = matchPronouns(wordSwap) // reflect pronouns
	return wordSwap                    // the word to be ammened to the response
}

/*
*	function to build the response
*	if the reponse contains "%"s" formatter
*	insert the wordSwap into the response
	else return response as it is
*/
func responseBuilder(response, wordSwap string) string {
	if strings.Contains(response, "%s") {
		return fmt.Sprintf(response, wordSwap)
	}
	return response
}

//AskEliza main function to start response generation
func AskEliza(input string) string {

	//create response[]response from file
	response := makeResponses("./database/responses.dat")
	randomResponse, _ := ReadLines("./database/randomResponses.dat")
	rand.Seed(time.Now().Unix())

	/*
	*	loop through repsonse struct
	*	if input conatins a pattern
	* 	Extract the main word
	*	Generate a random response for that pattern
	*	build a response and substitute the found word
	*	return response
	 */

	for _, response := range response {
		if response.Patterns.MatchString(input) {
			wordSwap := wordSwapper(response.Patterns, input)
			genResp := response.Answers[rand.Intn(len(response.Answers))]
			genResp = responseBuilder(genResp, wordSwap)
			return genResp
		}
	}
	// if no pattern match is found choose a random response
	return randomResponse[rand.Intn(len(randomResponse))]
}
