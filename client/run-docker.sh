docker run --net=host --rm \
    -e DISPLAY=unix$DISPLAY \
    -e MEMEFY_SERVER=gomano.de:8080 \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    --device /dev/snd \
    memefy_client
