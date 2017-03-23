# hksoundtouch

SoundTouch homekit accessory for [HomeControl](https://github.com/brutella/hc)

## Sample Code for using with HC

```golang
package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "github.com/llun/hksoundtouch"
)

func getSoundTouchAccessories() []*accessory.Accessory {
  services := soundtouch.Lookup()

  accessories := make([]*accessory.Accessory, len(services))
  for idx, service := range services {
    accessories[idx] = service.Accessory
  }

  return accessories
}

func main() {
  accessories := getSoundTouchAccessories()

  t, err := hc.NewIPTransport(hc.Config{
    Pin:  "32191123",
    Port: "51827",
  }, accessories[0], accessories[1:]...)
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
