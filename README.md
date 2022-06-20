# stumble-bot
Stumble Bot for Push Rank

## Getting Started

- Go language
- Git

## Setup

```bash
git clone https://github.com/ArugaZ/stumble-bot.git
cd stumble-bot
go build -ldflags "-s -w" main.go
```
1. Download App HTTPCanary (for sniff http request)
2. Run HTTPCanary and play stumble, take 1st winner at game (for boost)
3. Copy cookie at httpcanary (example at auth.txt)
## Usuage
Default :
- Paste cookie in auth.txt file
- round type is 0

Round Type :
- Check at httpcanary (example "http://kitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2/0") behind the url
1. 0 Mean you lost early
2. 1 Mean you lost
3. 2 Mean you lost at 3rd
4. 3 Mean you are the winner 

Default

```bash
./main stumble
```
Custom
```bash
./main stumble -a "Paste Your Auth" -r "Round Type"
```
Example with default auth
```bash
./main stumble -r 2
```

## TODO
- Nanti kl gabut setelah bab

## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/ArugaZ/stumble-bot/blob/main/README.md) for more details
