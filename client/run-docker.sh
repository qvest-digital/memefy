docker run --net=host --rm \
    -e DISPLAY=unix$DISPLAY \
    -e MEMEFY_SERVER=localhost:8080 \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    memefy_client
