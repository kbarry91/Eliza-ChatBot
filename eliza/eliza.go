package eliza

import (
	"regexp"
)

/*
type Response struct {
	pattern string
	answers []string
}

create []Response from file.
*/

// loop throug file and create a list of Responses
// file, err = os.Open(path)
// bufio.NewScanner(file), .Scan() .Text()

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
func AskEliza(input string) string {
	// send it in
	// process regex
	// do complicated stuff
	// return answer

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
