package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//create struct of responses
type Response struct {
	//Pattern string
	//pattern of respnonse
	Patterns *regexp.Regexp
	//answer for that response
	Answers []string
}

/*
func Response builder()[]Response{

}
// load file
*/

//function to make responses and populate struct
func makeResponses(path string) []Response {
	fullFile, _ := ReadLines(path)   //read in slice of all lines in the file.
	responses := make([]Response, 0) // make a slice of Responses to hold the responses, don't know how many there will be so start at size = 0
	for i := 0; i < len(fullFile); i += 2 {
		allPatterns := strings.Split(fullFile[i], ";")    // patterns on first line
		allResponses := strings.Split(fullFile[i+1], ";") // responses on the next line.
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern              // make pattern case insensitive.
			Patterns := regexp.MustCompile(pattern) // throws an error if the pattern doesn't compile.
			responses = append(responses, Response{Patterns: Patterns, Answers: allResponses})
		}
	}
	return responses
}

//function to test if responses are being populated
func PrintResponses(path string) {
	response := makeResponses(path)
	fmt.Printf("%+v\n", response)
}

//function to read all lines from file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//Response.Pattern = scanner.Text()
	}
	return lines, scanner.Err()
}

/*
//create []Response from file.

// loop throug file and create a list of Responses
//file, err = os.Open(path)
//bufio.NewScanner(file), .Scan() .Text()

/*
take user input

for resp in every response
	if resp.pattern matches user input
		extract from user input
		pick random answer
		sub from input into answer
		return answer
else
	return "generic answer"

*/
/*
func getResponses() []Response {
	raw, err := ioutil.ReadFile("./database/respJson.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Response
	json.Unmarshal(raw, &c)
	return c
}
func (r Response) toString() string {
	return toJson(r)
}

func toJson(r interface{}) string {
	bytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
func BuildResponses() int {
	responses := getResponses()
	for _, r := range responses {
		fmt.Println(r.toString())
	}

	fmt.Println(toJson(responses))
	//test print struct
	for _, t := range Response {
		fmt.print(Response.Pattern)
		fmt.print(Response.Answers)
	}
	return 1
} //buildResponses
*/
func AskEliza(input string) string {

	// send it in
	// process regex
	// do complicated stuff
	// return answer

	//create response[]response from file
	response := makeResponses("./database/responses.dat")
	rand.Seed(time.Now().UTC().UnixNano())
	for _, response := range response {
		if response.Patterns.MatchString(input) {
			//genResp := response.Answers[rand.Intn(len("Answers"))]
			//genResp := response.Answers[rand.Intn(len("response.Answers"))]
			fmt.Print("randnum " + strconv.Itoa(rand.Intn(len("response.Answers"))))
			genResp := response.Answers[0]

			return genResp
		}
	}

	pattern := "name is (.*)"
	likePattern := "I like (.*)"

	re := regexp.MustCompile(pattern)

	if re.MatchString(input) { // the input and regular expression match.
		match := re.FindStringSubmatch(input)
		name := match[1]
		return "Hello " + name + " it's nice to meet you."
	}
	re = regexp.MustCompile(likePattern)
	if re.MatchString(input) { // the input and regular expression match.
		match := re.FindStringSubmatch(input)
		like := match[1]
		return "Why do you like " + like + "?"
	}

	return "And how does that make you feel?"
}
