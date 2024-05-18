# **clx**
## Generate CLI commands for common tasks, the easy way!

**clx** takes the powerful functionality from [Sh4yy](https://github.com/Sh4yy)'s [gist](https://gist.github.com/Sh4yy/3941bf5014bc8c980fad797d85149b65) and makes it simple to use as a command-line tool.

## Installation

Running the following command will download and install the latest version of **clx**:

```bash
curl -sSL https://raw.githubusercontent.com/The-Daishogun/clx/main/install.sh | bash
```
## Usage

**clx** streamlines creating commands for common tasks. Simply provide a clear description of what you want to achieve, and clx will generate the corresponding command!

**Getting started:**

1.  **First Use:** On your first run, clx will guide you through a one-time configuration process to set your preferences. This includes selecting the desired model and setting your [Groq](https://console.groq.com/) API token.
2.  **Generating Commands:** Once configured, use `clx` followed by a description of the task you want to accomplish. 

**Example:**

```bash
$ clx kill process running on port 8000
```

This will generate the appropriate command to terminate a process using port 8000.

## More Examples:

Here are some additional examples to showcase clx's versatility:

* **Network Tasks:**
  * `clx ping google.com`
  * `clx scan for open ports on 192.168.1.100`
* **File Management:**
  * `clx create new directory called documents`
  * `clx search for all files containing "error" in current directory`
  * `clx convert all jpg images to png in this folder`
* **System Administration:**
  * `clx list all running processes`
  * `clx check system uptime`
  * `clx update all packages`

**Remember, the clearer and more specific your description, the more accurate the generated command will be!**

## Configuration

**clx** stores your configuration in `~/.config/clx/config.json`. This file manages your preferences, such as the selected model and API token. feel free to change it to try out different models or swap out your api token.

## Contributing

We welcome contributions to **clx**! If you have improvements or new functionalities in mind, feel free to fork the repository on GitHub and submit a pull request.

**Happy automating with clx!**