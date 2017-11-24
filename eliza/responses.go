// Author: Matthew Shiel
// Code adapted from https://github.com/kennysong/goeliza/

package eliza

// Adapted from https://github.com/kennysong/goeliza

// Introductions is a list of introduction sentences for ELIZA.
var Introductions = []string{
	"Hello, How are you feeling today?",
	"How do you do. Are you seeking help today?",
	"Please tell me what's been bothering you.",
	"Is something troubling you?",
}

// Goodbyes is a list of goodbye sentences for ELIZA.
var Goodbyes = []string{
	"Farewell. It was lovely speaking with you.",
	"Thank you for talking with me today.",
	"Thank you, that will be $150. Have a good day!",
	"Goodbye. This was nice, hopefully we do it again sometime.",
	"Goodbye. I'm looking forward to our next session.",
	"Well.. I guess time is up, call back anytime!",
	"Maybe we could discuss this over more in our next session? Goodbye.",
	"Ciao",
}

// Psychobabble may be slightly non-deterministic, since map iteration may be out
// of order, so a broader regex may be matched before a more specific one.
var Psychobabble = map[string][]string{
	`i need (.*)`: {
		"Why do you need %s?",
		"Would it really help you to get %s?",
		"Are you sure you need %s?",
	},
	`why don'?t you ([^\?]*)\??`: {
		"Do you really think I don't %s?",
		"Perhaps I will %s eventually.",
		"Do you really want me to %s?",
	},
	`why can'?t I ([^\?]*)\??`: {
		"Do you think you should be able to %s?",
		"If you could %s, what would you do?",
		"I don't know -- why can't you %s?",
		"Have you really tried?",
	},
	`^eliza`: {
		"That's the name, therapy is my game.",
		"Yes?",
		"That's me.",
		"Ah so you remember my name.",
	},
	`(.*) your name?`: {
		"My name is Eliza",
		"I thought that would be self explanatory.",
		"....Look up.",
	},
	`i can'?t (.*)`: {
		"How do you know you can't %s?",
		"Perhaps you could %s if you tried something else?",
		"What would it take for you to %s?",
	},
	`i am (.*)`: {
		"Did you come to me because you are %s?",
		"How long have you been %s?",
		"Why don't you tell me why you're %s?",		
		"How do you feel about being %s?",
		"Why do you think you're %s?",		
	},
	`i'?m (.*)`: {
		"How does being %s really make you feel?",
		"Hm, why are you %s?",
		"Why don't you tell me why you're %s?",
		"Why do you think you're %s?",
	},
	`(my name is|my name's) (.*)`: {
		"Hello %s how are you today?",
		"What's up %s?",
		"Hello, how are you feeling today?",
		"Hello, Is something troubling you?",
	},
	`are you ([^\?]*)\??`: {
		"Why does it matter whether I am %s?",
		"Would you prefer it if I were not %s?",
		"Perhaps you believe I am %s.",
		"I may be %s -- what do you think?",
		"If you choose to believe so but I'd rather talk about you.",
	},
	`what (.*)`: {
		"Why do you ask?",
		"How would an answer to that help you?",
		"Could you phrase that a little differently please?",
		"What do you mean?",
	},
	`how (.*)`: {
		"How do you suppose?",
		"Perhaps you can answer your own question.",
		"What is it you're really asking?",
		"Let's think about that, what are you really asking?",
	},
	`because (.*)`: {
		"Is that the real reason?",
		"What other reasons come to mind?",
		"Does that reason apply to anything else?",
		"If %s, what else must be true?",
	},
	`(.*) (sorry|apologies) (.*)`: {
		"There's no need to apologize, I'm here for you.",
		"You've done nothing wrong, why are you sorry?",
	},
	`^hello(.*)`: {
		"Hello... I'm glad you could drop by today.",
		"Hi there... how are you today?",
		"Hello, how are you feeling today?",
		"Hello, Is something troubling you?",
	},
	`^hi(.*)`: {
		"Hello... I'm glad you could drop by today.",
		"Hi there... how are you today?",
		"Hello, how are you feeling today?",
		"Hello, Is something troubling you?",
	},
	`^hey(.*)`: {
		"Hello... I'm glad you could drop by today.",
		"Hi there... how are you today?",
		"Hello, how are you feeling today?",
		"Hello, Is something troubling you?",
	},
	`^yo(.*)`: {
		"Hello... I'm glad you could drop by today.",
		"Hi there... how are you today?",
		"Hello, how are you feeling today?",
		"Hello, Is something troubling you?",
	},
	`^thanks(.*)`: {
		"You're welcome!",
		"Anytime!",
		"You're more than welcome",
	},
	`^thank you(.*)`: {
		"You're welcome",
		"Anytime!",
		"You're more than welcome",
	},
	`^good morning(.*)`: {
		"Good morning... I'm glad you could drop by today.",
		"Good morning... how are you today?",
		"Good morning, how are you feeling today?",
	},
	`^good afternoon(.*)`: {
		"Good afternoon... I'm glad you could drop by today.",
		"Good afternoon... how are you today?",
		"Good afternoon, how are you feeling today?",
	},
	`I think (.*)`: {
		"Do you doubt %s?",
		"Do you really think so?",
		"But you're not sure %s?",
	},
	`(.*) friend (.*)`: {
		"Tell me more about your friends.",
		"When you think of a friend, what comes to mind?",
		"Why don't you tell me about a childhood friend?",
	},
	`^(?:yes|yeah|affirmitive|yup)$`: {
		"You seem quite sure.",
		"OK, but can you elaborate a bit?",
		"Are you sure?",
		"Is there something else that's troubling you?",
	},
	// Non capturing group
	`^(?:no|nope|nah)$`: {
		"Are you just being negative or do you mean that?",
		"Could you provide a little more than just 'no'?",
		"Are you sure?",
		"Is there something else that's troubling you?",
		"Never short of words I see..",
		"What else is on your mind?",
	},
	`(.*) ago`: {
		"I see and has much changed since then?",
		"Hm, I see and how do you feel about that?",
		"Why don't you tell me how this started?",
	},
	`(.*) computer(.*)`: {
		"Are you really talking about me?",
		"Does it seem strange to talk to a computer?",
		"How do computers make you feel?",
		"Do you feel threatened by computers?",
	},
	`is it (.*)`: {
		"Do you think it is %s?",
		"Perhaps it's %s -- what do you think?",
		"If it were %s, what would you do?",
		"It could well be that %s.",
	},
	`(.*) i guess (.*)`: {
		"You guess?",
		"So are you not sure?",
		"You don't sound very sure about that",
		"You sound hesitant, are you sure abou that?",
		"Do you really believe that yourself?",
	},
	`my (.*) was (.*)`: {
		"Your %s is %s?",
		"Did I hear that correctly, your %s is %s?",
		"How do you feel about that?",
	},
	`it is (.*)`: {
		"You seem very certain.",
		"If I told you that it probably isn't %s, what would you feel?",
	},
	`can you ([^\?]*)\??`: {
		"What makes you think I can't %s?",
		"If I could %s, then what?",
		"Why do you ask if I can %s?",
	},
	`(.*)dream(.*)`: {
		"Tell me more about your dream.",
	},
	`can i ([^\?]*)\??`: {
		"Perhaps you don't want to %s.",
		"Do you want to be able to %s?",
		"If you could %s, would you?",
	},
	`you are (.*)`: {
		"Why do you think I am %s?",
		"Does it please you to think that I'm %s?",
		"Perhaps you would like me to be %s.",
		"Perhaps you're really talking about yourself?",
	},
	`you'?re (.*)`: {
		"Why do you say I'm %s?",
		"Why do you think I am %s?",
		"Are we talking about you, or me?",
		"Haha very funny..",
	},
	`i don'?t (.*)`: {
		"Don't you really %s?",
		"Why don't you %s?",
		"Do you want to %s?",
	},
	`i feel (.*)`: {
		"Good, tell me more about these feelings.",
		"Do you often feel %s?",
		"When do you usually feel %s?",
		"When you feel %s, what do you do?",
	},
	`i (hate|loathe|detest|dislike|despise) (.*)`: {
		"Why do you %s %s?",
		"Are you just saying you %s %s?",
		"How long have you been feeling this way?",
		"I'm sorry to hear that, when did these feelings begin?",
	},
	`i have (.*)`: {
		"Why do you tell me that you've %s?",
		"Have you really %s?",
		"Now that you have %s, what will you do next?",
	},
	`i would (.*)`: {
		"Could you explain why you would %s?",
		"Why would you %s?",
		"Who else knows that you would %s?",
	},
	`is there (.*)`: {
		"Do you think there is %s?",
		"It's likely that there is %s.",
		"Would you like there to be %s?",
	},
	`my (.*)`: {
		"I see, your %s.",
		"So what's wrong with your %s?",
		"Your %s, how do you feel?",
	},
	`you (.*)`: {
		"We should be discussing you, not me.",
		"We're talking about you not me remember?",
		"Look, I'd rather that we talk about you",
	},
	`why (.*)`: {
		"Why don't you tell me the reason why %s?",
		"Why do you think %s?",
	},
	`i want (.*)`: {
		"What would it mean to you if you got %s?",
		"Why do you want %s?",
		"What would you do if you got %s?",
		"If you got %s, then what would you do?",
	},
	`(.*) mother(.*)`: {
		"Tell me more about your mother.",
		"What was your relationship with your mother like?",
		"How do you feel about your mother?",
		"How does this relate to your feelings today?",
		"Good family relations are important.",
	},
	`(.*) father(.*)`: {
		"Tell me more about your father.",
		"How did your father make you feel?",
		"How do you feel about your father?",
		"Does your relationship with your father relate to your feelings today?",
		"Do you have trouble showing affection with your family?",
	},
	`(.*) child(.*)`: {
		"Did you have close friends as a child?",
		"What is your favorite childhood memory?",
		"Do you remember any dreams or nightmares from childhood?",
		"Did the other children sometimes tease you?",
		"How do you think your childhood experiences relate to your feelings today?",
	},
	`(.*)\?`: {
		"Why do you ask that?",
		"Please consider whether you can answer your own question.",
		"Perhaps the answer lies within yourself?",
		"Why don't you tell me?",
	},
	`\b(thank you|thanks)\b`: {
		"You're welcome!",
		"Always here to help",
		"Anytime",
		"What else can I do to help you?",
	},
}

// DefaultResponses are called If ELIZA doesn't understand the question, then it will reply with one of
// these default responses
var DefaultResponses = []string{
	"Please tell me more.",
	"We could discuss that but why don't we start with your family?",
	"Tell me, how is your relationship with your father?",
	"Do you get along with your mother?",
	"Can you elaborate on that?",
	"I see. Please go on.",
	"Very interesting... ",
	"I see. And what does that tell you?",
	"How does that make you feel?",
	"How do you feel when you say that?",
}

// QuitResponses contains  list of statements that indicate the user wants to end the conversation
var QuitResponses = []string{
	"goodbye",
	"bye",
	"quit",
	"exit",
}

// ReflectedWords is a table to reflect words in question fragments inside the response.
// eg. "User: Are you real? Eliza: I am real"
var ReflectedWords = map[string]string{
	"am":     "are",
	"was":    "were",
	"i":      "you",
	"i'd":    "you would",
	"i've":   "you have",
	"i'll":   "you will",
	"my":     "your",
	"are":    "am",
	"you've": "I have",
	"you'll": "I will",
	"your":   "my",
	"yours":  "mine",
	"you":    "me",
	"me":     "you",
}
