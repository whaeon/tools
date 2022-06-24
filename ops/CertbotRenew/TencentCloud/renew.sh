#!/bin/bash

# absolute path
FilePath=`pwd`

echo $FilePath
go build -o $FilePath/renew $FilePath/main.go
export SecretID=""
export SecretKey=""

# get main domain
num=$(echo $CERTBOT_DOMAIN | awk -F "." '{print NF}')
# Outside loop loop determines whether it is the main domain
for i in `seq $num`
do
    # generate filter condition
    condition=""
    for ((c=$i;c<=$num;c++))
    do
        val=`printf "\$%s"$c`
        if [ $c -eq $num ];then
            condition=$condition$val
        else
            condition=$condition$val,
        fi
    done
    # loop to get domains
    region=$(echo $CERTBOT_DOMAIN | awk -F "." "{print $condition}" OFS=".")
    # loop to Determine whether the domain name is the main domain
    res=$(whois $region | grep Expiration)
    if [ ! -z "$res" ];then
        # get subdomain
        condition=""
        for ((n=1;n<$i;n++))
        do
            val=`printf "\$%s"$n`
            max=`expr $i - 1`
            if [ $n -eq $max ];then
                condition=$condition$val
            else
                condition=$condition$val,
            fi
        done
        subDomain=$(echo $CERTBOT_DOMAIN | awk -F "." "{print $condition}" OFS=".")
        break
    fi
done

SubDomainForTXTRecord=_acme-challenge.$subDomain

$FilePath/renew -D $region -d $SubDomainForTXTRecord -t "TXT" -v $CERTBOT_VALIDATION -l "默认"

sleep 30