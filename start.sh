ffmpeg -re -i h264 -codec:v libx264  -g 100  -f hls -hls_list_size 3 -hls_wrap 10 -hls_time 10 -hls_init_time 10 playlist.m3u8

