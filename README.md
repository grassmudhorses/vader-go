# vader-go
Go port of the [VADER sentiment analysis tool](https://github.com/cjhutto/vaderSentiment), with various improvements.

For full information, please read the [VADER-Sentiment-Analysis README](https://github.com/cjhutto/vaderSentiment).

### Source

Based on research from
> Hutto, C.J. & Gilbert, E.E. (2014). VADER: A Parsimonious Rule-based Model for Sentiment Analysis of Social Media Text. Eighth International Conference on Weblogs and Social Media (ICWSM-14). Ann Arbor, MI, June 2014.

## Usage

```go
import "github.com/grassmudhorses/vader-go/sentitext"
...
vader := sentitext.PolarityScore("Hello World 💕 I Love the World!")
```
## Development

1. https://golang.org/doc/install
2. ```go get github.com/grassmudhorses/vader-go/sentitext```
3. ```cd $GOPATH/src/github.com/grassmudhorses/vader-go```
4. ```go test ./...```

## Deployment

After following Development above:
1. Follow https://cloud.google.com/gcp/getting-started/ to create a new Gooogle Cloud Project
2. Install the gcloud tool https://cloud.google.com/sdk/install
3. Log in ```gcloud auth login```
4. ```gcloud functions deploy vader --runtime go111 --trigger-http --entry-point GoogleCloudFunctionHTTP --project [YOUR PROJECT NAME] --allow-unauthenticated```
5. Go to the URL specified in the output: ```httpsTrigger: ```\
```url: https://[YOURLOCATION].cloudfunctions.net/vader```
6. GET
    ```http
    https://[YOURLOCATION].cloudfunctions.net/vader/Hello World 💕 I Love the World!
    ```

    Result:
    ```json
    {"neg":0,"neu":0.5718097779472029,"pos":0.4281902220527971,"compound":0.6697392619941973}
    ```

## Roadmap

1. Full feature parity with the original VADER-Sentiment-Analysis python tool (90% complete)
2. Allow users to load lexicon information from any source (80% complete)
3. Code cleanup and golang-specific optimizations (50% complete)
4. Iterative undefined word analysis using Wiktionary and Urbandictionary (0% complete)
5. Lexicon generation through iterative dictionary analysis (0% complete)