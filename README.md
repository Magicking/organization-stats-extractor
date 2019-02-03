# Organisation stats extractor

Generate screenshot of [Code frequency](https://github.com/Magicking/quorum-stress/graphs/code-frequency)
for each public repository in the `repo_url.json` using [ChromeDP](https://github.com/chromedp/chromedp).

## Usage

```sh
cat > repo_url.json <<HEREDOC
[
 "Magicking/organization-stats-extractor"
]
HEREDOC
go run main.go
```
