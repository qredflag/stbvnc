# stbvnc
Quick-and-dirty remote access to STB Qt/GUI via framebuffer.


##How to build
```shell script
sudo apt update
sudo apt install golang
cd ~/projects/stbvnc/src
GOOS=linux GOARCH=arm go build -o stbvnc
```

##How to run
Just execute stbvnc. index.html should reside in same folder.

```shell script
./stbvnc
```

###To do
* Capturing of HiSilicon video surface data
* Encoding screen image to  JPEG or H.264 with hardware acceleration, instead of RAW blobs.     
* Security with server/client certificates (self-signed?!)
* Full-featured mouse or/and keyboard support.
* Audio casting.