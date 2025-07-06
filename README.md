```
         ▄▄▄▄▄                                                              
       ▄▀  ▄█▀▀ ▀▀▄        ▐██▀▀██ ███▀██     ██▀▀██▌███▀██ ███▀██ ▀██▀██▀
     █▀█ ▐▌ ▄█▀▀█▄▐▌       ▐██  ▀▀ ██▌ ██     ██  ▀▀ ██▌ ██ ▐██ ██▌ ██  ▀
    █  █ ▐█▀ ▀█▄▄ ▀█       ▐██ ▄█▄ ██▌▌██     ██     ██▌▌██ ▐██ ██▌ ████
    ▀█ ▀▀█▌   █▌ █  █      ▐██ ▐██ ██▌ ██     ██  ▄█▌██▌ ██ ▐██ ██▌ ██ ▄▄
     ▐▀▀▄▄▄█▀▀▐▌ █▄█▀      ▐██▄███ ███▄██     ██▄▄██ ███▄██ ▐██▄██▌▄██▄██▌
      █▄   ▄▄▀▀ ▄█
        ▀▀▀▀█▄▄▀▀
```

# GO-CODE: AI Coding Agent for the Terminal

GO-CODE is a terminal-based AI coding agent. It helps you automate coding tasks, execute shell commands, and manage files directly from your terminal. GO-CODE is designed for developers who want to interact with their codebase and operating system efficiently using natural language.

## Features

- Execute shell commands securely from the terminal
- Read, write, and delete files and directories
- Search for patterns in files with `grep`
- Visualize project structure with `tree` and `ls`
- Maintain conversational context and history
- Thorough logging of all actions and AI responses
- Tool-augmented LLM responses for autonomous codebase navigation
- Azure OpenAI integration for enterprise deployments

## Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/KacemMathlouthi/go-code.git
   cd go-code
   ```

2. **Set up your environment variables**  
   Create a `.env` file or export the following variables:
   ```
   AZURE_API_KEY=your-azure-api-key
   AZURE_ENDPOINT=your-azure-endpoint
   AZURE_API_VERSION=your-api-version
   AZURE_DEPLOYMENT_NAME=your-deployment-name
   ```

3. **Build and run**
   ```bash
   go build -o go-code
   ./go-code
   ```

4. **Interact with GO-CODE**
   - Type your coding or shell-related requests.
   - Use `--help` for available commands, `--quit` to exit, and `--clear` to reset conversation history.

## Example Usage

```shell
$ ./go-code
> Create a React app, make a TO-DO app, with scheduling, deadlines, similar to Google
Calendar, install dependencies, and launch the app
> Make the page more beautiful, make it dark mode, and add more features like task search and more... 
```

## License

MIT License

---

© 2025 Kacem Mathlouthi
