# hksoundtouch

SoundTouch homekit accessory for [HomeControl](https://github.com/brutella/hc)

## Sample Code for using with HC

```golang
package main

import (
  "github.com/brutella/hc"
  "github.com/llun/hksoundtouch"
)

func main() {
  soundtouch := soundtouch.NewAccessory()
  t, err := hc.NewIPTransport(hc.Config{
    Pin:  "32191123",
    Port: "51827",
  }, soundtouch.Accessory)
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    t.Stop()
  })

  t.Start()
}
```

# License

MIT
