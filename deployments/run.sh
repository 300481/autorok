#!/bin/bash
docker run -it --network host --cap-add=NET_ADMIN -e "AUTOROK_CONFIG_URL=$AUTOROK_CONFIG_URL" -e "S6_KEEP_ENV=1" autorok