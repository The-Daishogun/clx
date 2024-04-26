# clx
Generate CLI commands for common tasks.


### General idea and most of the code is from [Sh4yy](https://gist.github.com/Sh4yy)'s [gist](https://gist.github.com/Sh4yy/3941bf5014bc8c980fad797d85149b65).
I just wanted a simple way to use the code so here we are.


## Installation
```bash
curl -sSL https://raw.githubusercontent.com/The-Daishogun/clx/main/install.sh --output - | sudo bash
```

## Usage
#### Setting the Api Token
you need to set your groq api token first.
head over to https://console.groq.com/keys and grab an api **token**. then run:

```bash
clx set-token <YOUR_API_TOKEN>
```

#### Actual usage
after that the usage is simple.


```bash
clx <PROMPT>
```

#### example usage:

```bash
$ clx kill process running on port 8000
```
