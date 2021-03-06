# Data Representation and Querying Project Eliza-ChatBot
My name is Kevin Barry I am a third year software developement student in G.M.I.T. 
This repository was created as part of a project for the module [Data Representation and Querying](https://data-representation.github.io/)

The objective of this project was to create a chatbot web application in Go. The project description can be found [here](https://data-representation.github.io/problems/project.html) 

#### What is an Eliza Chatbot
ELIZA is an early natural language processing computer program created from 1964 to 1966 at the MIT Artificial Intelligence Laboratory by Joseph Weizenbaum. In this project the aim was to create a chatbot that resembled a human speaking bot and give the illusion that the machine was learning and engaging in conversation.
To find out more about the Eliza program here is a link to which I refered to in order to gauge a more indept understanding of what Eliza is all about [Wikipedia ELIZA](https://en.wikipedia.org/wiki/ELIZA)

#### How to run the program
This program uses the Go programming language .If you do not currently have Go installed click on the following link to download [INSTALL GO](https://golang.org/dl/)

To clone the repository to your local machine, in command prompt enter 
```
git clone https://github.com/kbarry91/Eliza-ChatBot.git
```
There is two ways to run this program
1. **Build and Run** Navigate to the Eliza-Chatbot folder and enter the following to compile the code 
```
go build chatbot.go
```
This will create a .exe file in your current directory.To launch the server and run the file 
```
chatbot
```
2. **Run** to simply run the program in your command prompt enter the following 
```
go run chatbot.go
```  
The webpage will now be served to the local host. To view the homepage ,open you browser and enter to the local address
```
127.0.0.1:8080
```
Or alternatively you can enter
```
localhost:8080
```
### The Developement process
#### Research
In order to succesfully complete this project I followed the basic prinicpal of *how do you get an elephant in a room?bit by bit*. The first step of this process was to research and gather information on the task ahead. To do this I used many online resources. I was also able to find versions of Eliza that other developers had implemented and learned how I could adapt certain algorithims and improve were possible. This gave a good idea of what defined a good or bad chat bot.

#### Design
I used the principal of abstraction to make developing this program as easy as possible. I started by handwriting pseudo-code to give me a general path of how the code would run. I decided to structure the program in a way that would allow methods to run independently which would make it easier to troubleshoot bugs along the way.

#### Implementing the algorithim into code 
The first step of coding the problem was to get the program to serve the HTML file to the localhost. Once I had this working I used ajax and jQuery to pass the user entered values from the HTML page to the server and back to the HTML page. I then built the eliza.go side as a seperate programme which contained all the computations to build a response that was simply called by a main method in chatbot.go. I decided to create an external responses.dat file to hold all the patterns and responses so that this can be easily extended to add more responses without the need to having to edit the .go code.

### Added Features
To add a more realisitic messenger feel to the program I added mp3 format audio files that play when a message is sent and recieved,turn the volume up to try this. I put a delay in the repsonse to give the illusion that the other person i.e. the chatbot , was typing out a message to reply. Just for fun I added a few responses to give life to Eliza such as telling a joke or singing a song. The home page is both modern and colourfull to attract the user with text dialog boxes,avatars and a constant scroll to bottom to ensure your most recent question is always in view

### References
I used many different sites in order to learn about Eliza and learn new algorithims that were needed to develop the chatbot. Any adapted code has been referenced within  the code.
Here is a list of a few resourcses I visited most frequently.
+ (https://stackoverflow.com/)
+ (https://www.w3schools.com/jquery/)
+ (http://api.jquery.com/jquery.ajax/)
+ (https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/)
+ (https://data-representation.github.io/)

### Screen Capture Of Program

![Screen capture of Eliza](/static/img/screencapture.jpg?raw=true "Screen Capture of Eliza")