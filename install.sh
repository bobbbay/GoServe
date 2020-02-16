binary=`uname -m`

git clone https://github.com/Bobbbay/GoServe.git && curl --output ./GoServe/GoServe "https://raw.githubusercontent.com/Bobbbay/GoServe/master/binaries/GoServe-${binary}" && cd ./GoServe && chmod u+x ./GoServe