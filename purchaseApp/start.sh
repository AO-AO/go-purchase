#! /bin/bash

CURDIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd )
echo $CURDIR 
cd $CURDIR
go build

if [[ "$REVIEW_ENV"x = "production"x ]]
then
    ./purchaseApp
else
    ./purchaseApp 
fi

