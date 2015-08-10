# flowpush
A simple GoLang lib to push messages to a flow in flowdock

## Example
import "github.com/mleone896/flowpush"
```
const(
    FlowdockAPIToken = <the flow you wish to push to>
    )
func main() {

    flowpush.PushMessageToFlowWithKey(FlowdockAPIToken, "message to push ", "user")
}

```

