# flowpush
A simple GoLang lib to push messages to a flow in flowdock

## Example
import "github.com/mleone896/flowpush"
```
func main() {
const(
FlowdockAPIToken = <the flow you wish to push to>

flowpush.PushMessageToFlowWithKey(FlowdockAPIToken, "it just worked", "ec2events")

```

