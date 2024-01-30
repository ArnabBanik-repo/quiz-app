# CLI Quiz Application


## Overview

This is a Command Line Interface (CLI) quiz application written in Go. The application allows users to take a quiz by providing a CSV file containing questions and answers. Users can also set a time limit for answering each question.

## Features

- Load quiz questions from a CSV file.
- Set a time limit for answering the quiz.
- Display questions and options to the user.
- Record and display the user's score at the end of the quiz.
- Support for various question types (e.g., multiple-choice, true/false).


## Installation

### Compile from source

#### Pre-requisites
- git
- go

```{bash}
mkdir quizApp
cd quizApp
git clone github.com/ArnabBanik-repo/quiz-app .
go build .
```

### Download the executable binary from the Releases page (Linux)


## Usage

```{bash}
./quizApp -csv <questions> -limit <time_in_seconds>
```

Example
```{bash}
./quizApp --file quiz_questions.csv --time 60
```

For help
```{bash}
./quizApp -h
```

## CSV File Format
```
question,option1,option2,option3,option4,correct_option,score
What is the capital of France?,Paris,Berlin,London,Madrid,Paris,5
```

Feel free to contribute, report issues, or suggest improvements!
