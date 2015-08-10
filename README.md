https://travis-ci.org/mleone896/flowpush.svg?branch=master
# flowpush
A simple GoLang lib to push messages to a flow in flowdock

## USAGE 

```
package main

import "github.com/mleone896/flowpush"

const(
    FlowdockAPIToken = <the flow you wish to push to>
    )
func main() {

    flowpush.PushMessageToFlowWithKey(FlowdockAPIToken, "message to push ", "user")
}

```

