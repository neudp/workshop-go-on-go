# Go on go Workshop
Here is code to conduct go no go workshop

## Disclaimer
ALL MATERIALS ARE INFORMATIVE. THEY ARE LICENSED AS PUBLIC DOMAIN, IF OTHER IS NOT DEFINED EXPLICITLY.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Topics

- [X] Hello World, Monoprocess application
- [X] Command line application
- [X] Environment variables
- [X] HTTP client  (check for best practices, FastHTTP)
- [X] HTTP server
- [X] Serialization
- [X] Handling HTTP requests
- [X] Dependency injection (vanilla deep) <-
- [X] Testing
- [X] Handling errors
- [ ] sql package
- [ ] ORM
- [ ] Websockets
- [ ] gRPC server and protobuf
- [ ] Concurrency
- [ ] Memory management
- [ ] go work

## Prerequisites
if you have go installed, you can use it. 
If you don't have go installed, you can use docker.

## Docker
```bash
docker compose up -d --build
```

## Run examples (Go)
```bash
go run ./cmd/01-hello-world
```

## Run examples (Docker)
```bash
./run.sh go run ./cmd/01-hello-world
```
