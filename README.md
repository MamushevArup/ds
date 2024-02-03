How to set up this discord bot. 
1) Go to discord for devs create bot and get token.
2) Paste token in .env file TOKEN=<token_here>
3) Clone this project and use cd discord command
4) Use docker compose up --build command that brings to you 3 service mongo, go-server, bot-handler
5) Use bot in your server with commands below

Available commands for discord-bot

ATTENTION: first of all use !help command to know how to use other commands.
1) !hello - command greet the bot and start the interactions
2) !help - command show all available command, descriptions and usage
3) !game - create a game where bot generate number
4) !guess - user try to find the number
5) !poll - command create a poll with options
6) !vote - user can vote on a specific poll and option

Idea explanation

Idea is keep a code modular and independent.
If you want to use server without discord bot you can use server endpoints
Implemented layered architecture approach.
Two dockerfile created for bot and server.
All other practices described in code. 
Thank you for attention!!!