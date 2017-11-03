#! /bin/bash

CURDIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd )
echo $CURDIR 
cd $CURDIR
go build

if [[ "$SERVICE_ENV"x = "production"x ]]
then
    ./purchaseApp &> log/out.log 1> log/info.log 2> log/error.log
else
    ./purchaseApp 
fi

