#!/bin/bash

export $(grep -v '^#' .env | xargs -d '\n')
