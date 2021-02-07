#!/bin/sh
(cd web && npm install && npm run build && rm build/static/**/*.map)
