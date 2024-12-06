package quote

import "math/rand"

func DefaultQuotesProvider() []Quote {
	quotesRaw := []map[string]string{
		{"quote": "The only way to do great work is to love what you do.", "author": "Steve Jobs"},
		{"quote": "In the middle of every difficulty lies opportunity.", "author": "Albert Einstein"},
		{"quote": "Success is not the key to happiness. Happiness is the key to success.", "author": "Albert Schweitzer"},
		{"quote": "Believe you can and you're halfway there.", "author": "Theodore Roosevelt"},
		{"quote": "Don't watch the clock; do what it does. Keep going.", "author": "Sam Levenson"},
		{"quote": "Success is not final, failure is not fatal: It is the courage to continue that counts.", "author": "Winston Churchill"},
		{"quote": "Hardships often prepare ordinary people for an extraordinary destiny.", "author": "C.S. Lewis"},
		{"quote": "What you do today can improve all your tomorrows.", "author": "Ralph Marston"},
		{"quote": "Strive not to be a success, but rather to be of value.", "author": "Albert Einstein"},
		{"quote": "The only limit to our realization of tomorrow is our doubts of today.", "author": "Franklin D. Roosevelt"},
		{"quote": "Your time is limited, so don't waste it living someone else's life.", "author": "Steve Jobs"},
		{"quote": "The best way to predict the future is to invent it.", "author": "Alan Kay"},
		{"quote": "It does not matter how slowly you go as long as you do not stop.", "author": "Confucius"},
		{"quote": "Success usually comes to those who are too busy to be looking for it.", "author": "Henry David Thoreau"},
		{"quote": "The future belongs to those who believe in the beauty of their dreams.", "author": "Eleanor Roosevelt"},
		{"quote": "Opportunities don't happen. You create them.", "author": "Chris Grosser"},
		{"quote": "Act as if what you do makes a difference. It does.", "author": "William James"},
		{"quote": "Work hard in silence, let your success be your noise.", "author": "Frank Ocean"},
		{"quote": "Success is walking from failure to failure with no loss of enthusiasm.", "author": "Winston Churchill"},
		{"quote": "Happiness is not something ready-made. It comes from your own actions.", "author": "Dalai Lama"},
		{"quote": "If you want to achieve greatness, stop asking for permission.", "author": "Anonymous"},
		{"quote": "The only way to achieve the impossible is to believe it is possible.", "author": "Charles Kingsleigh"},
		{"quote": "The best revenge is massive success.", "author": "Frank Sinatra"},
		{"quote": "Do what you can, with what you have, where you are.", "author": "Theodore Roosevelt"},
		{"quote": "The purpose of our lives is to be happy.", "author": "Dalai Lama"},
		{"quote": "Courage is resistance to fear, mastery of fear—not absence of fear.", "author": "Mark Twain"},
		{"quote": "You miss 100 percent of the shots you don't take.", "author": "Wayne Gretzky"},
		{"quote": "Do one thing every day that scares you.", "author": "Eleanor Roosevelt"},
		{"quote": "Dream big and dare to fail.", "author": "Norman Vaughan"},
		{"quote": "What lies behind us and what lies before us are tiny matters compared to what lies within us.", "author": "Ralph Waldo Emerson"},
		{"quote": "Success is how high you bounce when you hit bottom.", "author": "George S. Patton"},
		{"quote": "Limitations live only in our minds. But if we use our imaginations, our possibilities become limitless.", "author": "Jamie Paolinetti"},
		{"quote": "Life is what happens when you're busy making other plans.", "author": "John Lennon"},
		{"quote": "The way to get started is to quit talking and begin doing.", "author": "Walt Disney"},
		{"quote": "The only thing standing between you and your goal is the story you keep telling yourself as to why you can't achieve it.", "author": "Jordan Belfort"},
		{"quote": "Success is the sum of small efforts, repeated day-in and day-out.", "author": "Robert Collier"},
		{"quote": "Keep your face always toward the sunshine—and shadows will fall behind you.", "author": "Walt Whitman"},
		{"quote": "Don't be pushed around by the fears in your mind. Be led by the dreams in your heart.", "author": "Roy T. Bennett"},
		{"quote": "Start where you are. Use what you have. Do what you can.", "author": "Arthur Ashe"},
		{"quote": "Failure is the condiment that gives success its flavor.", "author": "Truman Capote"},
		{"quote": "It always seems impossible until it's done.", "author": "Nelson Mandela"},
		{"quote": "Opportunities multiply as they are seized.", "author": "Sun Tzu"},
		{"quote": "Fall seven times and stand up eight.", "author": "Japanese Proverb"},
		{"quote": "Don't let yesterday take up too much of today.", "author": "Will Rogers"},
		{"quote": "Challenges are what make life interesting and overcoming them is what makes life meaningful.", "author": "Joshua J. Marine"},
		{"quote": "The only person you are destined to become is the person you decide to be.", "author": "Ralph Waldo Emerson"},
		{"quote": "I am not a product of my circumstances. I am a product of my decisions.", "author": "Stephen R. Covey"},
		{"quote": "Do not go where the path may lead, go instead where there is no path and leave a trail.", "author": "Ralph Waldo Emerson"},
		{"quote": "Everything you've ever wanted is on the other side of fear.", "author": "George Addair"},
		{"quote": "Whether you think you can or you think you can't, you're right.", "author": "Henry Ford"},
	}
	quotes := make([]Quote, 0)

	for _, value := range quotesRaw {
		quotes = append(quotes, Quote{
			Author: value["author"],
			Text:   value["quote"],
		})
	}

	return quotes
}

func GenerateSingleRandomQuote(quotesProvider func() []Quote) Quote {
	quotes := quotesProvider()
	randomIndex := rand.Intn(len(quotes) - 1)

	return quotes[randomIndex]
}
