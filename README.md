# VoIP Caller by 1vere$k

This is simple 5060 port exploit that calls to a receiver VoIP device with pre-defined phrase.  
The app establishes a TCP connection and write a VoIP Payload with the INVITE method which initiates a call with pre-defined "caller-name" which is our `phrase`.

## Usage

Application

```
1. git clone https://github.com/iveresk/voip-caller.git
2. cd voip-caller
3. go build -o /voip-caller
4. ./voip-caller -t <target_url or file_name> [OPTIONAL] -d false -s <YOUR_PHRASE>

- t - target URL or filename
- d - debug {true, false} flag which enables or disables logs
- s - is your "caller name" which is your phrase
```

Docker with pre-defined 20k ruzzian VoIP devices.

```
docker run -d masterroot/voip-caller
```

Have fun with `ruzzland-terrorist-state` VoIP devices :D


## Example
[![Zelensky Calling. We Will Find You!](https://res.cloudinary.com/marcomontalbano/image/upload/v1657031991/video_to_markdown/images/vimeo--727047847-c05b58ac6eb4c4700831b2b3070cd403.jpg)](https://vimeo.com/727047847 "Zelensky Calling. We Will Find You!")

## RUSSIAN WARSHIP GO FRAK YOURSELF!
![Russian Warship Go F*ck Yourself](rus-ship-go-fuck-yourself.jpeg)

## Contact
You are free to contact me via [Keybase](https://keybase.io/1veresk) for any details. 