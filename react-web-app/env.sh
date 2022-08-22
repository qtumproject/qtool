#!/bin/bash

# Recreate config file
rm -rf ./env-config.js
touch ./env-config.js

# Add assignment 
echo "window._env_ = {" >> ./env-config.js
varvalue="${REACT_APP_API_URL:-}"
varname="REACT_APP_API_URL"
echo "  $varname: \"$varvalue\"," >> ./env-config.js
echo "}" >> ./env-config.js