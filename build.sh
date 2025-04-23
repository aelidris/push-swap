#!/bin/bash

cd Checker
go build -o ../checker .
cd ../pushswap
go build -o ../push-swap .
cd ..
