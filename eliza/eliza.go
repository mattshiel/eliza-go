package eliza

// Code adapted from https://github.com/kennysong/goeliza/

import (
    "fmt"
    "math/rand"
    "regexp"
    "strings"
)

// Greetings will return a random introductory sentence for ELIZA.
func Greetings() string {
    return randChoice(Introductions)
}

// GoodbyeResponse will return a random goodbye sentence for ELIZA.
func GoodbyeResponse() string {
    return randChoice(Goodbyes)
}

// ReplyTo will construct a reply for a given statement using ELIZA's rules.
func ReplyTo(statement string) string {
    // First, preprocess the statement for more effective matching
    statement = preprocess(statement)

    // Then, we check if this is a quit statement
    if IsQuitStatement(statement) {
        return GoodbyeResponse()
    }

    // Next, we try to match the statement to a statement that ELIZA can 
    // recognize, and construct a pre-determined, appropriate response.
    for pattern, responses := range Psychobabble {
        re := regexp.MustCompile(pattern)
        matches := re.FindStringSubmatch(statement)

        // If the statement matched any recognizable statements.
        if len(matches) > 0 {
            // If we matched a regex group in parentheses, get the first match.
            // The matched regex group will match a "fragment" that will form 
            // part of the response, for added realism.
            var fragment string
            if len(matches) > 1 {
                fragment = reflect(matches[1])
            }

            // Choose a random appropriate response, and format it with the 
            // fragment, if needed.
            response := randChoice(responses)
            if strings.Contains(response, "%s") {
                response = fmt.Sprintf(response, fragment)
            }
            return response
        }
    }

    // If no patterns were matched, return a default response.
    return randChoice(DefaultResponses)
}

// IsQuitStatement returns if the statement is a quit statement
func IsQuitStatement(statement string) bool {
    statement = preprocess(statement)
    for _, quitResponse := range QuitResponses {
        if statement == quitResponse {
            return true
        }
    }
    return false
}

// preprocess will do some normalization on a statement for better regex matching
func preprocess(statement string) string {
    statement = strings.TrimRight(statement, "\n.!")
    statement = strings.ToLower(statement)
    return statement
}

// reflect flips a few words in an input fragment (such as "I" -> "you").
func reflect(fragment string) string {
    words := strings.Split(fragment, " ")
    for i, word := range words {
        if reflectedWord, ok := ReflectedWords[word]; ok {
            words[i] = reflectedWord
        }
    }
    return strings.Join(words, " ")
}

// randChoice returns a random element in an (string) array.
func randChoice(list []string) string {
    randIndex := rand.Intn(len(list))
    return list[randIndex]
}