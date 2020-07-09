# gobot-servo

Servo motor controller for Raspberry Pi.

It's using [gobot](https://gobot.io/) and requires pi-blaster(https://github.com/sarfata/pi-blaster) for PWM. Discussion of pi-blaster is [here](https://github.com/hybridgroup/gobot/issues/691).

## Usage

Execute a server on Raspberry Pi.

```
$ go install
$ ./gobot-servo &
```

Change the servo angle from API.

```
$ curl localhost:8080/open
$ curl localhost:8080/close
```
